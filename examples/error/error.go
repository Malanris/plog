package main

import (
	"fmt"

	"github.com/Malanris/plog"
)

func main() {
	err := fmt.Errorf("too much sugar")
	log.Error("failed to bake cookies", "err", err)
}
