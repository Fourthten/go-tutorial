package main

import (
	"go-rest/banking/app"
	"go-rest/banking/logger"
	// "log"
)

func main() {

	// log.Println("Starting our application")
	logger.Info("Starting our application")
	app.Start()

}
