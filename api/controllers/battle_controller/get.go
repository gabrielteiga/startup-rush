package battle_controller

import (
	"net/http"
	"path"
	"strconv"

	"github.com/gabrielteiga/startup-rush/api/responses"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/battle_entity"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/event_entity"
	"github.com/gabrielteiga/startup-rush/internal/domain/services"
)

const GET_SUCCESSFULLY_MESSAGE string = "Battle has been listed successfully"
const GET_ERROR_MESSAGE string = "Something went wrong with the Battle listing"

type GetBattleById struct {
	tournamentService *services.TournamentService
}

func NewGetBattleByID(tournamentService *services.TournamentService) *GetBattleById {
	return &GetBattleById{
		tournamentService: tournamentService,
	}
}

func (gb *GetBattleById) Handle(w http.ResponseWriter, r *http.Request) {
	idString := path.Base(r.URL.Path)
	battleID, err := strconv.Atoi(idString)
	if err != nil {
		responses.JSONErrorResponse[any](w, http.StatusInternalServerError, GET_ERROR_MESSAGE, nil)
		return
	}

	battle, err := gb.tournamentService.GetBattleByID(uint(battleID))
	if err != nil {
		responses.JSONErrorResponse[any](w, http.StatusInternalServerError, GET_ERROR_MESSAGE, nil)
		return
	}

	events, err := gb.tournamentService.GetEvents()
	if err != nil {
		responses.JSONErrorResponse[any](w, http.StatusInternalServerError, GET_ERROR_MESSAGE, nil)
		return
	}

	fullBattle := &struct {
		Battle *battle_entity.Battle `json:"battle"`
		Events []*event_entity.Event `json:"events"`
	}{
		Battle: battle,
		Events: events,
	}

	responses.JSONSuccessResponse(w, http.StatusOK, GET_SUCCESSFULLY_MESSAGE, &fullBattle)
}
