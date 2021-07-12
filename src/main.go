package main

import (
	"github.com/dbielecki97/bookstore-items-api/src/app"
	"github.com/dbielecki97/bookstore-utils-go/logger"
)

func main() {
	logger.Info("Starting Items API server...")
	app.StartApplication()
}
