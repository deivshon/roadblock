package main

import (
	"main/pkg/globalconfig"
	"main/pkg/utils"
	"os"
	"path/filepath"
)

func getAllConfigs(rootConfigDir string, globalConfigPath string) (globalconfig.GlobalConfig, []string) {
	configDir, configDirErr := os.UserConfigDir()

	if rootConfigDir == "" {
		if configDirErr != nil {
			utils.Log.Fatalf("could not get user configuration directory: %v", configDirErr)
		}

		rootConfigDir = filepath.Join(configDir, roadblockConfigDir)
	}

	globalConfigPassedAsArg := globalConfigPath != ""
	if globalConfigPath == "" {
		if configDirErr != nil {
			utils.Log.Fatalf("could not get user configuration directory: %v", configDirErr)
		}

		globalConfigPath = filepath.Join(configDir, roadblockConfigDir, roadblockGlobalConfigPath)
	}

	configPaths, err := utils.SearchFilesRecursive(rootConfigDir, configsName)
	if err != nil {
		utils.Log.Fatalf("could not recursively search configurations root directory: %v", err)
	}

	globalConfig, err := globalconfig.GetGlobalConfig(globalConfigPath, globalConfigPassedAsArg)
	if err != nil {
		utils.Log.Fatalf("could not get global roadblock configuration: %v", err)
	}

	return globalConfig, configPaths
}
