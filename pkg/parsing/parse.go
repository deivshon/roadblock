package parsing

import (
	"fmt"
	"slices"

	"github.com/mattn/go-shellwords"
)

var DefaultWrappers = []string{
	"sudo",
	"doas",
	"time",
	"strace",
	"nice",
}

func ParseCommand(rawCommand string, opts CommandParsingOpts) ([]string, error) {
	command, err := shellwords.Parse(rawCommand)
	if err != nil {
		return nil, fmt.Errorf("could not parse command into shellwords: %v", err)
	}

	wrappers := opts.Wrappers
	if opts.AddDefaultWrappers == nil || *opts.AddDefaultWrappers {
		for _, w := range DefaultWrappers {
			if !slices.Contains(wrappers, w) {
				wrappers = append(wrappers, w)
			}
		}
	}

	if opts.RemoveWrappers == nil || *opts.RemoveWrappers {
		command = removeWrappers(command, wrappers)
	}

	return command, nil
}

func removeWrappers(command []string, wrappers []string) []string {
	for len(command) > 0 && slices.Contains(wrappers, command[0]) {
		command = slices.Delete(command, 0, 1)
	}

	return command
}
