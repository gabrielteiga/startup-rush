package startup_controller

import (
	"net/http"
	"path"
	"strconv"

	"github.com/gabrielteiga/startup-rush/api/responses"
	"github.com/gabrielteiga/startup-rush/internal/domain/services"
)

const GET_SUCCESSFULLY_MESSAGE string = "Participants has been listed successfully"
const GET_ERROR_MESSAGE string = "Something went wrong with the Participants listing"

type GetStartupByTournamentID struct {
	TournamentService *services.TournamentService
	StartupService    *services.StartupService
}

func NewGetStartupsByTournamentID(tournamentService *services.TournamentService, startupService *services.StartupService) *GetStartupByTournamentID {
	return &GetStartupByTournamentID{
		TournamentService: tournamentService,
		StartupService:    startupService,
	}
}

func (ct *GetStartupByTournamentID) Handle(w http.ResponseWriter, r *http.Request) {
	idString := path.Base(r.URL.Path)
	TournamentID, err := strconv.Atoi(idString)
	if err != nil {
		responses.JSONErrorResponse[any](w, http.StatusInternalServerError, GET_ERROR_MESSAGE, nil)
		return
	}

	participants := ct.TournamentService.FindParticipantsByTournamentID(uint(TournamentID))
	if participants == nil {
		responses.JSONErrorResponse(w, http.StatusInternalServerError, GET_ERROR_MESSAGE, &participants)
		return
	}

	var startupsIds []uint
	for _, participant := range participants {
		startupsIds = append(startupsIds, participant.StartupID)
	}

	startups := ct.StartupService.FindByIDs(startupsIds)
	if startups == nil {
		responses.JSONErrorResponse[any](w, http.StatusInternalServerError, GET_ERROR_MESSAGE, nil)
		return
	}
	responses.JSONSuccessResponse(w, http.StatusOK, GET_SUCCESSFULLY_MESSAGE, &startups)
}
