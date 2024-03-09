package utils

import (
	"log"
	"os"
)

var Log = log.New(os.Stderr, "roadblock: ", 0)
