package handlers

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Error(code int, message string) APIError {
	return APIError{
		Status:  code,
		Message: message,
	}
}

func SendAPIError(w http.ResponseWriter, code int, message string) {
	json.NewEncoder(w).Encode(APIError{
		Status:  code,
		Message: message,
	})
}
