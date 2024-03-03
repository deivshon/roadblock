package condition

import "fmt"

func (c *Condition) ValidationCheck() error {
	if len(c.Id) == 0 {
		return fmt.Errorf("ids must not be empty")
	}

	if err := c.Selector.ValidationCheck(); err != nil {
		return fmt.Errorf("selector error: %w", err)
	}

	if err := c.Evaluator.ValidationCheck(); err != nil {
		return fmt.Errorf("evaluator error: %w", err)
	}

	return nil
}
