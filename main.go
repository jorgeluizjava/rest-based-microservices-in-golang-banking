package main

import (
	"github.com/jorgeluizjava/banking/app"
	"github.com/jorgeluizjava/banking/app/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}

// DB_USER=root DB_PASSWD=root DB_ADDR=localhost DB_PORT=3306 DB_NAME=banking go run main.go
