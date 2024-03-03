package main

import (
	"main/pkg/config"
	"main/pkg/utils"
)

func main() {
	runtimeConfig, err := config.GetConfig("./examples/config.yaml")
	if err != nil {
		utils.Log.Fatalf("Could not get config: %v", err)
	}

	utils.Log.Printf("%+v", runtimeConfig)
}
