package condition

import (
	"errors"
	"fmt"
)

func (s *ConditionSelector) ValidationCheck() error {
	activeCount := 0

	if s.Command != nil {
		activeCount++
	}
	if s.WordIndex != nil {
		activeCount++
		if *s.WordIndex < 0 {
			return fmt.Errorf("received bad word index `%v`, must be positive", *s.WordIndex)
		}
	}
	if s.AnyWord != nil {
		activeCount++
	}
	if s.EveryWord != nil {
		activeCount++
	}

	if activeCount != 1 {
		return errors.New("multiple selectors active")
	}

	return nil
}
