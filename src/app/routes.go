package app

import (
	"github.com/dbielecki97/bookstore-items-api/src/controller"
	"net/http"
)

func createUrlMappings() {
	router.HandleFunc("/items", controller.ItemController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{item_id}", controller.ItemController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items/{item_id}", controller.ItemController.Update).Methods(http.MethodPut)
	router.HandleFunc("/ping", controller.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items/search", controller.ItemController.Search).Methods(http.MethodPost)
	router.HandleFunc("/items/{item_id}", controller.ItemController.Delete).Methods(http.MethodDelete)
}
