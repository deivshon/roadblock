package main

import (
	"flag"
	"main/pkg/config"
	"main/pkg/utils"

	"github.com/mattn/go-shellwords"
)

func main() {
	command := flag.String("t", "", "Target command to evaluate")

	flag.Parse()

	runtimeConfig, err := config.GetConfig("./examples/config.yaml")
	if err != nil {
		utils.Log.Fatalf("could not get config: %v", err)
	}

	parsedCommand, err := shellwords.Parse(*command)
	if err != nil {
		utils.Log.Fatalf("could not parse command into shellwords: %v", err)
	}

	forbidden, err := runtimeConfig.IsCommandForbidden(parsedCommand)
	if err != nil {
		utils.Log.Fatalf("could not evaluate command: %v", err)
	}

	if forbidden.Forbidden {
		utils.Log.Fatalf("forbidden by rule `%v`", forbidden.BrokenRule)
	}
}
