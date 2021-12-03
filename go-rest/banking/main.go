package main

import (
	"go-rest/banking-lib/logger"
	"go-rest/banking/app"
	// "log"
)

func main() {

	// log.Println("Starting our application")
	logger.Info("Starting our application")
	app.Start()

}
