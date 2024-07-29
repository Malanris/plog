package main

import "github.com/Malanris/plog"

func startOven(degree int) {
	log.Helper()
	log.Info("Starting oven", "degree", degree)
}

func main() {
	log.SetReportCaller(true)
	startOven(400)
}
