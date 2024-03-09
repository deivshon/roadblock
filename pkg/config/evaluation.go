package config

import (
	"fmt"
	"main/pkg/condition"
)

func (r *RoadblockConfig) IsCommandForbidden(command []string) (ForbiddenResult, error) {
	conditionsMap := make(map[string]condition.Condition)
	for _, cond := range r.Conditions {
		_, exists := conditionsMap[cond.Id]
		if exists {
			return ForbiddenResult{}, fmt.Errorf("multiple conditions have `%v` as their id", cond.Id)
		}

		conditionsMap[cond.Id] = cond
	}

	for _, rule := range r.Rules {
		forbidden, err := rule.Evaluate(func(condId string) (bool, error) {
			condition, exists := conditionsMap[condId]
			if !exists {
				return false, fmt.Errorf("condition `%v` not found", condId)
			}

			return condition.Evaluate(command)
		})
		if err != nil {
			return ForbiddenResult{}, fmt.Errorf("could not evaluate rule `%v`: %w", rule.Name, err)
		}

		if forbidden {
			return ForbiddenResult{
				Forbidden:  true,
				BrokenRule: rule.Name,
			}, nil
		}
	}

	return ForbiddenResult{
		Forbidden: false,
	}, nil
}
