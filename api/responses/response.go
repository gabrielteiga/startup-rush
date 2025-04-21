package responses

import (
	"encoding/json"
	"net/http"
)

const SUCCESS_STATUS string = "Success"
const ERROR_STATUS string = "Error"

type Response[T any] struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    *T     `json:"data,omitempty"`
}

func JSONSuccessResponse[T any](w http.ResponseWriter, httpCode int, message string, data *T) error {
	return createJsonResponse(w, httpCode, message, SUCCESS_STATUS, data)
}

func JSONErrorResponse[T any](w http.ResponseWriter, httpCode int, message string, data *T) error {
	return createJsonResponse(w, httpCode, message, ERROR_STATUS, data)
}

func createJsonResponse[T any](w http.ResponseWriter, httpCode int, message, status string, data *T) error {
	response := NewReponse(message, status, data)

	w.WriteHeader(httpCode)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(response)
}

func NewReponse[T any](message, status string, data *T) *Response[T] {
	return &Response[T]{
		Message: message,
		Status:  status,
		Data:    data,
	}
}
