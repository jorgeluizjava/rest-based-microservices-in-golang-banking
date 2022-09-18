package main

import (
	"github.com/jorgeluizjava/banking/app"
	"github.com/jorgeluizjava/banking/app/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
