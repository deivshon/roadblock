package condition

import (
	"errors"
	"fmt"
	"strings"
)

func (s *ConditionSelector) evaluateSelection(command []string, isForbidden func(string) (bool, error)) (bool, error) {
	if s.Command != nil && *s.Command {
		wholeCommand := strings.Join(command, " ")

		res, err := isForbidden(wholeCommand)
		if err != nil {
			return false, err
		}

		return res, nil
	}
	if s.WordIndex != nil {
		idx := *s.WordIndex
		if idx >= len(command) {
			return false, nil
		}
		if idx < 0 {
			return false, fmt.Errorf("word index `%v` is not positive", idx)
		}

		res, err := isForbidden(command[idx])
		if err != nil {
			return false, err
		}

		return res, nil
	}
	if s.AnyWord != nil && *s.AnyWord {
		for _, word := range command {
			res, err := isForbidden(word)
			if err != nil {
				return false, err
			}

			if res {
				return true, nil
			}
		}

		return false, nil
	}
	if s.EveryWord != nil && *s.EveryWord {
		for _, word := range command {
			res, err := isForbidden(word)
			if err != nil {
				return false, err
			}

			if !res {
				return false, nil
			}
		}

		return true, nil
	}

	return false, errors.New("no selector active")
}

func (s *ConditionSelector) validationCheck() error {
	activeCount := 0

	if s.Command != nil && *s.Command {
		activeCount++
	}
	if s.WordIndex != nil {
		activeCount++
		if *s.WordIndex < 0 {
			return fmt.Errorf("received bad word index `%v`, must be positive", *s.WordIndex)
		}
	}
	if s.AnyWord != nil && *s.AnyWord {
		activeCount++
	}
	if s.EveryWord != nil && *s.EveryWord {
		activeCount++
	}

	if activeCount == 0 {
		return errors.New("no selector active")
	}
	if activeCount != 1 {
		return errors.New("multiple selectors active")
	}

	return nil
}
