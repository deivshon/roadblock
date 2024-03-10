package globalconfig

import (
	"bytes"
	"errors"
	"fmt"
	"main/pkg/parsing"
	"os"

	"gopkg.in/yaml.v3"
)

var t = true
var defaultGlobalConfig = GlobalConfig{
	Parsing: parsing.CommandParsingOpts{
		AddDefaultWrappers: &t,
		RemoveWrappers:     &t,
	},
}

func GetGlobalConfig(path string, errorOnInexistant bool) (GlobalConfig, error) {
	pathInfo, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) && !errorOnInexistant {
			return defaultGlobalConfig, nil
		}

		return GlobalConfig{}, fmt.Errorf("could not stat `%v`: %w", path, err)
	}

	if pathInfo.IsDir() {
		return GlobalConfig{}, fmt.Errorf("global configuration path `%v` is a directory", path)
	}

	rawConfig, err := os.ReadFile(path)
	if err != nil {
		return GlobalConfig{}, fmt.Errorf("error reading file: %w", err)
	}

	var config GlobalConfig
	decoder := yaml.NewDecoder(bytes.NewReader(rawConfig))
	decoder.KnownFields(true)
	if err = decoder.Decode(&config); err != nil {
		return GlobalConfig{}, fmt.Errorf("error parsing: %w", err)
	}

	return config, nil
}
