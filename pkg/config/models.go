package config

import (
	"main/pkg/condition"
	"main/pkg/rule"
)

type RoadblockConfig struct {
	Conditions []condition.Condition `yaml:"conditions"`
	Rules      []rule.RuleWithMeta   `yaml:"rules"`
}

type ForbiddenResult struct {
	Forbidden  bool
	BrokenRule string
}
