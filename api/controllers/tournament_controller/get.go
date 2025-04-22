package tournament_controller

import (
	"net/http"
	"path"
	"strconv"

	"github.com/gabrielteiga/startup-rush/api/responses"
	"github.com/gabrielteiga/startup-rush/internal/domain/services"
)

const GET_SUCCESSFULLY_MESSAGE string = "Tournament has been listed successfully"
const GET_ERROR_MESSAGE string = "Something went wrong with the tournament listing"

type GetTournamentByID struct {
	TournamentService *services.TournamentService
}

func NewGetTournamentByID(tournamentService *services.TournamentService) *GetTournamentByID {
	return &GetTournamentByID{
		TournamentService: tournamentService,
	}
}

func (ct *GetTournamentByID) Handle(w http.ResponseWriter, r *http.Request) {
	idString := path.Base(r.URL.Path)
	TournamentID, err := strconv.Atoi(idString)
	if err != nil {
		responses.JSONErrorResponse[any](w, http.StatusInternalServerError, GET_ERROR_MESSAGE, nil)
		return
	}

	tournaments := ct.TournamentService.GetByID(uint(TournamentID))
	if tournaments == nil {
		responses.JSONErrorResponse(w, http.StatusInternalServerError, GET_ERROR_MESSAGE, &tournaments)
		return
	}

	responses.JSONSuccessResponse(w, http.StatusOK, GET_SUCCESSFULLY_MESSAGE, &tournaments)
}
