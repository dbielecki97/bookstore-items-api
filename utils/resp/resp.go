package resp

import (
	"encoding/json"
	"github.com/dbielecki97/bookstore-utils-go/errs"
	"github.com/dbielecki97/bookstore-utils-go/logger"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	encodeErr := json.NewEncoder(w).Encode(body)
	if encodeErr != nil {
		logger.Error("could not encode", encodeErr)
	}
}

func Error(w http.ResponseWriter, err *errs.RestErr) {
	JSON(w, err.StatusCode, err)
}
