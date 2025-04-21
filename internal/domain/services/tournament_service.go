package services

import (
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/startup_entity"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/tournament_entity"
)

type TournamentService struct {
	TournamentRepository tournament_entity.ITournamentRepository
	StartupRepository    startup_entity.IStartupRepository
}

func NewTournamentService(tournamentRepository tournament_entity.ITournamentRepository, startupRepository startup_entity.IStartupRepository) *TournamentService {
	return &TournamentService{
		TournamentRepository: tournamentRepository,
		StartupRepository:    startupRepository,
	}
}

func (ts *TournamentService) Create(startupIDs []uint) *tournament_entity.Tournament {
	startups := ts.StartupRepository.FindByIDs(startupIDs)
	if len(startups) != len(startupIDs) {
		return nil
	}

	return ts.TournamentRepository.Create(startups)
}
