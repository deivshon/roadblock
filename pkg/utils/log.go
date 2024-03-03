package utils

import (
	"log"
	"os"
)

var Log = log.New(os.Stderr, "", 0)
