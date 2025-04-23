package tournament_controller

import (
	"net/http"
	"path"
	"strconv"

	"github.com/gabrielteiga/startup-rush/api/responses"
	"github.com/gabrielteiga/startup-rush/internal/domain/services"
)

const RANKING_SUCCESSFULLY_MESSAGE string = "Ranking has been listed successfully"
const RANKING_ERROR_MESSAGE string = "Something went wrong with the Ranking listing"

type RankingByTournamentID struct {
	TournamentService *services.TournamentService
}

func NewGetRankingByTournamentID(tournamentService *services.TournamentService) *RankingByTournamentID {
	return &RankingByTournamentID{
		TournamentService: tournamentService,
	}
}
func (rt *RankingByTournamentID) Handle(w http.ResponseWriter, r *http.Request) {
	idString := path.Base(r.URL.Path)
	tournamentID, err := strconv.Atoi(idString)
	if err != nil {
		responses.JSONErrorResponse[any](w, http.StatusInternalServerError, RANKING_ERROR_MESSAGE, nil)
		return
	}

	ranking, err := rt.TournamentService.GetRanking(uint(tournamentID))
	if err != nil {
		responses.JSONErrorResponse[any](w, http.StatusInternalServerError, RANKING_ERROR_MESSAGE, nil)
		return
	}

	responses.JSONSuccessResponse(w, http.StatusOK, RANKING_SUCCESSFULLY_MESSAGE, &ranking)
}
