package main

import (
	"flag"
	"main/pkg/checker"
	"main/pkg/utils"
	"os"
	"path/filepath"
	"runtime"

	"github.com/mattn/go-shellwords"
)

const configsName = "config.yml"
const roadblockConfigDir = "roadblock"

const roadblockSkipEnv = "ROADBLOCK_SKIP"

func main() {
	if os.Getenv(roadblockSkipEnv) != "" {
		return
	}

	command := flag.String("t", "", "Target command to evaluate")
	rootConfigDir := flag.String("c", "", "Configurations root directory")
	showBrokenRuleConfig := flag.Bool("s", false, "When a command is forbidden, show the configuration path containing the broken rule")

	flag.Parse()

	if *rootConfigDir == "" {
		configDir, err := os.UserConfigDir()
		if err != nil {
			utils.Log.Fatalf("could not get user configuration directory: %v", err)
		}

		*rootConfigDir = filepath.Join(configDir, roadblockConfigDir)
	}

	configPaths, err := utils.SearchFilesRecursive(*rootConfigDir, configsName)
	if err != nil {
		utils.Log.Fatalf("could not recursively search configurations root directory: %v", err)
	}

	parsedCommand, err := shellwords.Parse(*command)
	if err != nil {
		utils.Log.Fatalf("could not parse command into shellwords: %v", err)
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
