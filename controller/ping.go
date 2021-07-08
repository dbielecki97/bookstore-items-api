package controller

import (
	"github.com/dbielecki97/bookstore-utils-go/logger"
	"net/http"
)

var (
	PingController pingController = &defaultPingController{}
)

const (
	pongResponse = "pong"
)

type pingController interface {
	Ping(http.ResponseWriter, *http.Request)
}

type defaultPingController struct {
}

func (d defaultPingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(pongResponse))
	if err != nil {
		logger.Error("could not write pong response", err)
	}
}
