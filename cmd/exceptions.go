package main

import (
	"fmt"
	"strings"
)

func commandIsException(command string) bool {
	bashPrefix := fmt.Sprintf("export %v=", roadblockSkipEnv)
	fishPrefix := fmt.Sprintf("set -gx %v ", roadblockSkipEnv)
	zshPrefix := bashPrefix

	prefixExceptions := []string{
		bashPrefix,
		fishPrefix,
		zshPrefix,
	}

	for _, prefix := range prefixExceptions {
		if strings.HasPrefix(command, prefix) {
			return true
		}
	}

	return false
}
