package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorDetail struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

func SendError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := ErrorResponse{
		Error: ErrorDetail{
			Message: message,
			Code:    statusCode,
		},
	}

	json.NewEncoder(w).Encode(resp)
}
