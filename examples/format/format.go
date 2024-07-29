package main

import (
	"fmt"
	"time"

	"github.com/Malanris/plog"
)

func main() {
	for item := 1; item <= 100; item++ {
		log.Info(fmt.Sprintf("Baking %d / 100 ...", item))
		time.Sleep(100 * time.Millisecond)
	}
}
