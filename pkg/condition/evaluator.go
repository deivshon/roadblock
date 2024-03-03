package condition

import "errors"

func (e *ConditionEvaluator) ValidationCheck() error {
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

	if activeCount != 1 {
		return errors.New("multiple evaluators active")
	}

	return nil
}
