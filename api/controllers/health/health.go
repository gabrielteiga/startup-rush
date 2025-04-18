package health

import (
	"encoding/json"
	"net/http"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

type ResponseHealth struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func (hc *HealthController) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body := &ResponseHealth{
		Message: "The app is healthy!",
		Status:  "Success",
	}

	if err := json.NewEncoder(w).Encode(body); err != nil {
		http.Error(w, "Error during the response creation", http.StatusInternalServerError)
	}
}
