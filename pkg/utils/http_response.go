package utils

import (
	"encoding/json"
	"net/http"

	"github.com/ghulammuzz/go-restful-template/internal/errors"
)

func SendError(w http.ResponseWriter, appErr *errors.AppError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errors.GetStatusCode(appErr.Code))
	json.NewEncoder(w).Encode(appErr)
}

func SendResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
