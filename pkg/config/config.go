package config

import (
	"bytes"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func GetConfig(configPath string) (RoadblockConfig, error) {
	rawConfig, err := os.ReadFile(configPath)
	if err != nil {
		return RoadblockConfig{}, fmt.Errorf("error reading file: %w", err)
	}

	var config RoadblockConfig
	decoder := yaml.NewDecoder(bytes.NewReader(rawConfig))
	decoder.KnownFields(true)
	if err = decoder.Decode(&config); err != nil {
		return RoadblockConfig{}, fmt.Errorf("error parsing: %w", err)
	}

	if err = config.ValidationCheck(); err != nil {
		return RoadblockConfig{}, fmt.Errorf("consistency error: %w", err)
	}

	return config, nil
}
