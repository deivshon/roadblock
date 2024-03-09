package rule

import (
	"errors"
)

func (r *RuleWithMeta) ValidationCheck() error {
	if len(r.Name) == 0 {
		return errors.New("a name is required")
	}

	return r.Rule.validationCheck()
}

func (r *Rule) validationCheck() error {
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

	if activeCount == 0 {
		return errors.New("found node with no operator active")
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
		if err := r.validationCheck(); err != nil {
			return err
		}
	}

	return nil
}
