package startup_controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gabrielteiga/startup-rush/api/requests"
	"github.com/gabrielteiga/startup-rush/api/responses"
	"github.com/gabrielteiga/startup-rush/internal/domain/services"
	"github.com/gabrielteiga/startup-rush/internal/utils/parsers"
)

const CREATED_SUCCESSFULLY_MESSAGE string = "Startup has been created successfully"
const CREATED_ERROR_MESSAGE string = "Something went wrong with the startup creation"

type CreateStartup struct {
	StartupService *services.StartupService
}

func NewCreateStartup(startupService *services.StartupService) *CreateStartup {
	return &CreateStartup{
		StartupService: startupService,
	}
}

func (sc *CreateStartup) Handle(w http.ResponseWriter, r *http.Request) {
	var request *requests.RequestStartupCreate

	json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()

	errValidation := requests.Validate(request)
	if errValidation != nil {
		responses.JSONErrorResponse(w, http.StatusBadRequest, CREATED_ERROR_MESSAGE, errValidation)
		return
	}

	foundationDate, err := parsers.StringDateToTime(request.Foundation)
	if err != nil {
		log.Fatalf("error during the date parse")
	}

	startup := sc.StartupService.Create(request.Name, request.Slogan, foundationDate)

	responses.JSONSuccessResponse(w, http.StatusCreated, CREATED_SUCCESSFULLY_MESSAGE, startup)
}
