package condition

import "fmt"

func (c *Condition) Evaluate(command []string) (bool, error) {
	evaluatorFunction, err := c.Evaluator.getEvaluator()
	if err != nil {
		return false, fmt.Errorf("could not get evaluator for condition `%v`: %w", c.Id, err)
	}

	return c.Selector.evaluateSelection(command, evaluatorFunction)

}

func (c *Condition) ValidationCheck() error {
	if len(c.Id) == 0 {
		return fmt.Errorf("selector has no id")
	}

	if err := c.Selector.validationCheck(); err != nil {
		return fmt.Errorf("selector error: %w", err)
	}

	if err := c.Evaluator.validationCheck(); err != nil {
		return fmt.Errorf("evaluator error: %w", err)
	}

	return nil
}
