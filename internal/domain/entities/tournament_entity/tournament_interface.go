package tournament_entity

import "github.com/gabrielteiga/startup-rush/internal/domain/entities/startup_entity"

type ITournamentRepository interface {
	Create(startups []*startup_entity.Startup) (*Tournament, error)
	List() ([]*Tournament, error)
	FindByID(id uint) (*Tournament, error)
	Finish(tournamentID uint, championID *uint) (*Tournament, error)
}
