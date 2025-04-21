package tournament_entity

import (
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/participations_entity"
)

type Tournament struct {
	ID           uint                                   `json:"id"`
	Finished     bool                                   `json:"finished"`
	ChampionID   *uint                                  `json:"championId,omitempty"`
	Participants []*participations_entity.Participation `json:"participants,omitempty"`
}

func NewTournament(id uint, finished bool, championId *uint, participants []*participations_entity.Participation) *Tournament {
	return &Tournament{
		ID:           id,
		Finished:     finished,
		ChampionID:   championId,
		Participants: participants,
	}
}
