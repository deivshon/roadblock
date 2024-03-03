package rule

import (
	"errors"
	"fmt"
)

func (r *Rule) ValidationCheck() error {
	return r.validationCheckAux(0)
}

func (r *Rule) validationCheckAux(recursionLevel int) error {
	if recursionLevel == 0 && (r.Name == nil || len(*r.Name) == 0) {
		return errors.New("all rules must specify a name")
	} else if recursionLevel != 0 && r.Name != nil {
		return fmt.Errorf("rejecting rule name `%v`: must be specified at root node of the rule", *r.Name)
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
		return errors.New("more than one operator active in the same node")
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
