package rule

import (
	"errors"
	"fmt"
)

func (r *RuleWithMeta) Evaluate(judge func(string) (bool, error)) (bool, error) {
	return r.Rule.evaluate(judge)
}

func (r *Rule) evaluate(judge func(string) (bool, error)) (bool, error) {
	if r.ConditionId != nil {
		res, err := judge(*r.ConditionId)
		if err != nil {
			return false, fmt.Errorf("error in judge function: %w", err)
		}

		return res, nil
	}
	if r.Not != nil {
		res, err := r.Not.evaluate(judge)
		if err != nil {
			return false, err
		}

		return !res, nil
	}
	if r.AllOf != nil {
		for _, rule := range *r.AllOf {
			res, err := rule.evaluate(judge)
			if err != nil {
				return false, err
			}
			if !res {
				return false, nil
			}
		}

		return true, nil
	}
	if r.OneOf != nil {
		for _, rule := range *r.OneOf {
			res, err := rule.evaluate(judge)
			if err != nil {
				return false, err
			}
			if res {
				return true, nil
			}
		}

		return false, nil
	}

	return false, errors.New("node has no active operator")
}
