package startup_entity

import (
	"time"

	"github.com/gabrielteiga/startup-rush/internal/domain/entities/participations_entity"
)

type Startup struct {
	ID             uint                                   `json:"id,omitempty"`
	Name           string                                 `json:"name"`
	Slogan         string                                 `json:"slogan"`
	Foundation     time.Time                              `json:"foundation"`
	Participations []*participations_entity.Participation `json:"participations,omitempty"`
}

func NewStartup(id uint, name, slogan string, foundation time.Time) *Startup {
	return &Startup{
		ID:         id,
		Name:       name,
		Slogan:     slogan,
		Foundation: foundation,
	}
}
