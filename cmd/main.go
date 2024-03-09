package main

import (
	"flag"
	"fmt"
	"main/pkg/config"
	"main/pkg/utils"
	"os"
	"path/filepath"

	"github.com/mattn/go-shellwords"
)

const configsName = "config.yml"
const roadblockConfigDir = "roadblock"

func main() {
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

	for _, p := range configPaths {
		runtimeConfig, err := config.GetConfig(p)
		if err != nil {
			utils.Log.Fatalf("could not get config `%v`: %v", p, err)
		}

		forbidden, err := runtimeConfig.IsCommandForbidden(parsedCommand)
		if err != nil {
			utils.Log.Fatalf("could not evaluate command: %v\nsource config: %v", err, p)
		}

		if forbidden.Forbidden {
			configErrorMessage := ""
			if *showBrokenRuleConfig {
				configErrorMessage = fmt.Sprintf("\nsource config: %v", p)
			}
			utils.Log.Fatalf("forbidden by rule `%v`%s", forbidden.BrokenRule, configErrorMessage)
		}
	}
}
