package health

import (
	"net/http"

	"github.com/gabrielteiga/startup-rush/api/responses"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (hc *HealthController) Handle(w http.ResponseWriter, r *http.Request) {
	if err := responses.JSONSuccessResponse[any](w, "The app is healthy!", nil); err != nil {
		http.Error(w, "Error during the response creation", http.StatusInternalServerError)
	}
}
