package main

import (
	"test/app"

	"test/config"
	"test/logger"
)

func main() {
	logger.Info("Starting application...")

	globalConfig := config.Global

	app.Run(globalConfig)
}
