package app

import (
	"github.com/dbielecki97/bookstore-items-api/controller"
	"net/http"
)

func createUrlMappings() {
	router.HandleFunc("/items", controller.ItemController.Create).Methods(http.MethodPost)
	router.HandleFunc("/ping", controller.PingController.Ping).Methods(http.MethodGet)
}
