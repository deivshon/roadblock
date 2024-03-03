package rule

import (
	"errors"
	"fmt"
)

func (r *Rule) ValidationCheck() error {
	return r.validationCheckAux(0)
}

func (r *Rule) validationCheckAux(recursionLevel int) error {
	if recursionLevel == 0 && (r.Message == nil || len(*r.Message) == 0) {
		return errors.New("all rules must specify a message to display when they match")
	} else if recursionLevel != 0 && r.Message != nil {
		return fmt.Errorf("rejecting error message `%v`: must be specified at root node of the rule", *r.Message)
	}

	activeCount := 0
	if r.AllOf != nil {
		activeCount++
	}
	if r.OneOf != nil {
		activeCount++
	}
	if r.Not != nil {
		activeCount++
	}
	if r.ConditionId != nil {
		activeCount++
	}

	if activeCount != 1 {
		return errors.New("more than one operator active in the same rule")
	}

	var children []Rule
	if r.AllOf != nil {
		children = *r.AllOf
	} else if r.OneOf != nil {
		children = *r.OneOf
	} else if r.Not != nil {
		children = append(children, *r.Not)
	} else if r.ConditionId != nil {
		return nil
	}

	for _, r := range children {
		err := r.validationCheckAux(recursionLevel + 1)
		if err != nil {
			return err
		}
	}

	return nil
}
