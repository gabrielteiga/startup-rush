package tournament_controller

import (
	"net/http"
	"path"
	"strconv"

	"github.com/gabrielteiga/startup-rush/api/responses"
	"github.com/gabrielteiga/startup-rush/internal/domain/services"
)

const START_SUCCESSFULLY_MESSAGE string = "Tournament has been started successfully"
const START_ERROR_MESSAGE string = "Something went wrong with the tournament starting"

type StartTournamentByID struct {
	TournamentService *services.TournamentService
}

func NewStartTournamentByID(tournamentService *services.TournamentService) *StartTournamentByID {
	return &StartTournamentByID{
		TournamentService: tournamentService,
	}
}

func (ct *StartTournamentByID) Handle(w http.ResponseWriter, r *http.Request) {
	idString := path.Base(r.URL.Path)
	TournamentID, err := strconv.Atoi(idString)
	if err != nil {
		responses.JSONErrorResponse[any](w, http.StatusInternalServerError, GET_ERROR_MESSAGE, nil)
		return
	}

	tournaments := ct.TournamentService.Start(uint(TournamentID))
	if tournaments == nil {
		responses.JSONErrorResponse(w, http.StatusInternalServerError, GET_ERROR_MESSAGE, &tournaments)
		return
	}

	responses.JSONSuccessResponse(w, http.StatusOK, GET_SUCCESSFULLY_MESSAGE, &tournaments)
}
