package checker

import (
	"fmt"
	"main/pkg/config"
)

type CheckerWorker struct {
	parsedCommand        []string
	passedChan           chan<- bool
	errorsChan           chan<- error
	brokenRulesChan      chan<- string
	inputPathsChan       <-chan string
	showBrokenRuleConfig bool
}

func NewCheckerWorker(parsedCommand []string, passedChan chan<- bool, errorsChan chan<- error, brokenRulesChan chan<- string, inputPathsChan <-chan string, showBrokenRuleConfig bool) *CheckerWorker {
	return &CheckerWorker{
		parsedCommand:        parsedCommand,
		passedChan:           passedChan,
		errorsChan:           errorsChan,
		brokenRulesChan:      brokenRulesChan,
		inputPathsChan:       inputPathsChan,
		showBrokenRuleConfig: showBrokenRuleConfig,
	}
}

func (w *CheckerWorker) StartWorker() {
	for p := range w.inputPathsChan {
		runtimeConfig, err := config.GetConfig(p)
		if err != nil {
			w.errorsChan <- fmt.Errorf("could not get config `%v`: %w", p, err)
		}

		forbidden, err := runtimeConfig.IsCommandForbidden(w.parsedCommand)
		if err != nil {
			w.errorsChan <- fmt.Errorf("could not evaluate command: %v\nsource config: %v", err, p)
		}

		if forbidden.Forbidden {
			configErrorMessage := ""
			if w.showBrokenRuleConfig {
				configErrorMessage = fmt.Sprintf("\nsource config: %v", p)
			}
			w.brokenRulesChan <- fmt.Sprintf("forbidden by rule `%v`%s", forbidden.BrokenRule, configErrorMessage)
		}

		w.passedChan <- true
	}
}
