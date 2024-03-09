package config

import (
	"fmt"
	"slices"
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

	if err := c.uniqueIdsCheck(); err != nil {
		return fmt.Errorf("ids are not unique: %w", err)
	}

	return nil
}

func (c *RoadblockConfig) uniqueIdsCheck() error {
	ids := make([]string, 0, len(c.Conditions))
	for _, c := range c.Conditions {
		if slices.Contains(ids, c.Id) {
			return fmt.Errorf("id `%v` belongs to multiple conditions", c.Id)
		}

		ids = append(ids, c.Id)
	}

	return nil
}
