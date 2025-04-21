package startup_controller

import (
	"net/http"

	"github.com/gabrielteiga/startup-rush/api/responses"
	"github.com/gabrielteiga/startup-rush/internal/domain/services"
)

const LIST_SUCCESSFULLY_MESSAGE string = "All startups could be found successfully"
const LIST_ERROR_MESSAGE string = "Something went wrong with the listing of startups"

type ListStartup struct {
	StartupService *services.StartupService
}

func NewListStartup(startupService *services.StartupService) *ListStartup {
	return &ListStartup{
		StartupService: startupService,
	}
}

func (sc *ListStartup) Handle(w http.ResponseWriter, r *http.Request) {
	startups := sc.StartupService.List()

	responses.JSONSuccessResponse(w, http.StatusOK, LIST_SUCCESSFULLY_MESSAGE, &startups)
}
