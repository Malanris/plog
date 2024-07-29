package main

import (
	"os"

	"github.com/Malanris/plog"
)

func main() {
	logger := log.New(os.Stderr)
	logger.Warn("chewy!", "butter", true)
}
