package main

import (
	"flag"
	"main/pkg/checker"
	"main/pkg/parsing"
	"main/pkg/utils"
	"os"
	"runtime"
)

const configsName = "config.yml"
const roadblockConfigDir = "roadblock"
const roadblockGlobalConfigPath = "roadblock.yml"

const roadblockSkipEnv = "ROADBLOCK_SKIP"

func main() {
	if os.Getenv(roadblockSkipEnv) != "" {
		return
	}

	command := flag.String("t", "", "Target command to evaluate")
	rootConfigDir := flag.String("c", "", "Configurations root directory")
	showBrokenRuleConfig := flag.Bool("s", false, "When a command is forbidden, show the configuration path containing the broken rule")
	globalConfigPath := flag.String("g", "", "Path to the global configuration")
	flag.Parse()

	if commandIsException(*command) {
		return
	}

	globalConfig, configPaths := getAllConfigs(*rootConfigDir, *globalConfigPath)

	parsedCommand, err := parsing.ParseCommand(*command, globalConfig.Parsing)
	if err != nil {
		utils.Log.Fatalf("could not parse command: %v", err)
	}

	errorsChan := make(chan error)
	passedChan := make(chan bool, 32)
	brokenRulesChan := make(chan string)
	inputPathsChan := make(chan string, len(configPaths))

	for range runtime.NumCPU() {
		w := checker.NewCheckerWorker(parsedCommand, passedChan, errorsChan, brokenRulesChan, inputPathsChan, *showBrokenRuleConfig)
		go w.StartWorker()
	}

	for _, p := range configPaths {
		inputPathsChan <- p
	}

	passedConfigs := 0
	for passedConfigs < len(configPaths) {
		select {
		case <-passedChan:
			passedConfigs++
		case msg := <-brokenRulesChan:
			utils.Log.Fatalf(msg)
		case err := <-errorsChan:
			utils.Log.Fatalf("%v", err)
		}
	}
}
