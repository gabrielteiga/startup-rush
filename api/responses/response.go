package responses

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    *T     `json:"data,omitempty"`
}

func JSONSuccessResponse[T any](w http.ResponseWriter, message string, data *T) error {
	return createJsonResponse(w, message, "Success", data)
}

func JSONErrorResponse[T any](w http.ResponseWriter, message string, data *T) error {
	return createJsonResponse(w, message, "Error", data)
}

func createJsonResponse[T any](w http.ResponseWriter, message, status string, data *T) error {
	response := newReponse(message, status, data)

	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(response)
}

func newReponse[T any](message, status string, data *T) *Response[T] {
	return &Response[T]{
		Message: message,
		Status:  status,
		Data:    data,
	}
}
