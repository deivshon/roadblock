package globalconfig

import "main/pkg/parsing"

type GlobalConfig struct {
	Parsing parsing.CommandParsingOpts `yaml:"parsing"`
}
