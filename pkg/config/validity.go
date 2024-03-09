package config

import (
	"fmt"
)

func (c *RoadblockConfig) ValidationCheck() error {
	for _, cond := range c.Conditions {
		if err := cond.ValidationCheck(); err != nil {
			return fmt.Errorf("error in condition `%v`: %w", cond.Id, err)
		}
	}

	for idx, rule := range c.Rules {
		if err := rule.ValidationCheck(); err != nil {
			return fmt.Errorf("error in rule #%v (`%v`): %w", idx+1, rule.Name, err)
		}
	}

	return nil
}
