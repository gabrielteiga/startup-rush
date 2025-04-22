package battle_controller

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"github.com/gabrielteiga/startup-rush/api/requests"
	"github.com/gabrielteiga/startup-rush/api/responses"
	"github.com/gabrielteiga/startup-rush/internal/domain/services"
)

const BATTLE_SUCCESSFULLY_MESSAGE string = "Battle has been listed successfully"
const BATTLE_ERROR_MESSAGE string = "Something went wrong with the Battle listing"

type BattleTournament struct {
	tournamentService *services.TournamentService
}

func NewBattleTournament(tournamentService *services.TournamentService) *BattleTournament {
	return &BattleTournament{
		tournamentService: tournamentService,
	}
}

func (gb *BattleTournament) Handle(w http.ResponseWriter, r *http.Request) {
	var request *requests.BattleRequest

	json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()

	errValidation := requests.Validate(request)
	if errValidation != nil {
		responses.JSONErrorResponse(w, http.StatusBadRequest, BATTLE_ERROR_MESSAGE, errValidation)
		return
	}

	idString := path.Base(r.URL.Path)
	battleID, err := strconv.Atoi(idString)
	if err != nil {
		responses.JSONErrorResponse[any](w, http.StatusInternalServerError, GET_ERROR_MESSAGE, nil)
		return
	}

	events := make(map[uint][]uint, len(request.Battle))
	for _, event := range request.Battle {
		events[event.StartupID] = event.EventIDs
	}

	something, err := gb.tournamentService.Battle(uint(battleID), events)
	if err != nil {
		responses.JSONErrorResponse[any](w, http.StatusInternalServerError, BATTLE_ERROR_MESSAGE, nil)
		return
	}

	responses.JSONSuccessResponse[any](w, http.StatusOK, BATTLE_SUCCESSFULLY_MESSAGE, &something)
}
