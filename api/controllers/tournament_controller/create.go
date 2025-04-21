package tournament_controller

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielteiga/startup-rush/api/requests"
	"github.com/gabrielteiga/startup-rush/api/responses"
	"github.com/gabrielteiga/startup-rush/internal/domain/services"
)

const CREATED_SUCCESSFULLY_MESSAGE string = "Tournament has been created successfully"
const CREATED_ERROR_MESSAGE string = "Something went wrong with the tournament creation"

type CreateTournament struct {
	TournamentService *services.TournamentService
}

func NewCreateTournament(tournamentService *services.TournamentService) *CreateTournament {
	return &CreateTournament{
		TournamentService: tournamentService,
	}
}

func (ct *CreateTournament) Handle(w http.ResponseWriter, r *http.Request) {
	var request *requests.RequestCreateTournament

	json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()

	errValidation := requests.Validate(request)
	if errValidation != nil {
		responses.JSONErrorResponse(w, http.StatusBadRequest, CREATED_ERROR_MESSAGE, errValidation)
		return
	}

	tournament := ct.TournamentService.Create(request.StartupsIDs)
	if tournament == nil {
		responses.JSONErrorResponse(w, http.StatusInternalServerError, CREATED_ERROR_MESSAGE, tournament)
		return
	}

	responses.JSONSuccessResponse(w, http.StatusCreated, CREATED_SUCCESSFULLY_MESSAGE, tournament)
}
