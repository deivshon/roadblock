package condition

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func (e *ConditionEvaluator) getEvaluator() (func(string) (bool, error), error) {
	if e.Equals != nil {
		return func(s string) (bool, error) {
			return s == *e.Equals, nil
		}, nil
	}
	if e.Contains != nil {
		return func(s string) (bool, error) {
			return strings.Contains(s, *e.Contains), nil
		}, nil
	}
	if e.Regex != nil {
		return func(s string) (bool, error) {
			res, err := regexp.MatchString(*e.Regex, s)
			if err != nil {
				return false, fmt.Errorf("could not evaluate regex: %w", err)
			}

			return res, nil
		}, nil
	}

	return nil, errors.New("no evaluator active")
}

func (e *ConditionEvaluator) validationCheck() error {
	activeCount := 0
	if e.Equals != nil {
		activeCount++
	}
	if e.Contains != nil {
		activeCount++
	}
	if e.Regex != nil {
		activeCount++
	}

	if activeCount == 0 {
		return errors.New("no evaluator active")
	}
	if activeCount != 1 {
		return errors.New("multiple evaluators active")
	}

	return nil
}
