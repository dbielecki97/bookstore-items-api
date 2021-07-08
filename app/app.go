package app

import (
	"github.com/dbielecki97/bookstore-utils-go/logger"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	createUrlMappings()

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8083",
	}

	logger.Fatal("could not start server", srv.ListenAndServe())
}
