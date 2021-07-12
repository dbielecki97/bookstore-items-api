package app

import (
	"github.com/dbielecki97/bookstore-items-api/src/client/es"
	"github.com/dbielecki97/bookstore-utils-go/logger"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	es.Init()

	createUrlMappings()

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8083",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.Fatal("could not start server", srv.ListenAndServe())
}
