package services

import (
	"time"

	"github.com/gabrielteiga/startup-rush/internal/domain/entities/startup_entity"
)

type StartupService struct {
	StartupRepository startup_entity.IStartupRepository
}

func NewStartupService(startupRepository startup_entity.IStartupRepository) *StartupService {
	return &StartupService{
		StartupRepository: startupRepository,
	}
}

func (ss *StartupService) Create(name, slogan string, foundation time.Time) *startup_entity.Startup {
	return ss.StartupRepository.Create(name, slogan, foundation)
}

func (ss *StartupService) List() []*startup_entity.Startup {
	startups := ss.StartupRepository.List()
	return startups
}
