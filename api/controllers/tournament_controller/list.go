package tournament_controller

import (
	"net/http"

	"github.com/gabrielteiga/startup-rush/api/responses"
	"github.com/gabrielteiga/startup-rush/internal/domain/services"
)

const LIST_SUCCESSFULLY_MESSAGE string = "Tournaments have been listed successfully"
const LIST_ERROR_MESSAGE string = "Something went wrong with the tournament listing"

type ListTournament struct {
	TournamentService *services.TournamentService
}

func NewListTournament(tournamentService *services.TournamentService) *ListTournament {
	return &ListTournament{
		TournamentService: tournamentService,
	}
}

func (ct *ListTournament) Handle(w http.ResponseWriter, r *http.Request) {
	tournaments := ct.TournamentService.List()
	if tournaments == nil {
		responses.JSONErrorResponse(w, http.StatusInternalServerError, LIST_ERROR_MESSAGE, &tournaments)
		return
	}

	responses.JSONSuccessResponse(w, http.StatusOK, LIST_SUCCESSFULLY_MESSAGE, &tournaments)
}
