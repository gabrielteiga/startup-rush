package gorm_adapter

import (
	"time"

	"github.com/gabrielteiga/startup-rush/database"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/startup_entity"
	"gorm.io/gorm"
)

type StartupGORMRepository struct {
	DB *gorm.DB
}

func NewStartupGORMRepository(db *gorm.DB) *StartupGORMRepository {
	return &StartupGORMRepository{
		DB: db,
	}
}

func (sr *StartupGORMRepository) Create(name, slogan string, foundation time.Time) *startup_entity.Startup {
	startup := database.Startup{
		Name:       name,
		Slogan:     slogan,
		Foundation: foundation,
	}

	err := sr.DB.Create(&startup).Error
	if err != nil {
		return nil
	}

	return startup_entity.NewStartup(startup.ID, startup.Name, startup.Slogan, startup.Foundation)
}

func (sr *StartupGORMRepository) List() []*startup_entity.Startup {
	var startups []database.Startup

	err := sr.DB.Select("id, name, slogan, foundation").Find(&startups).Error
	if err != nil {
		return nil
	}

	var startupEntities []*startup_entity.Startup
	for _, startupModel := range startups {
		startup := startup_entity.NewStartup(startupModel.ID, startupModel.Name, startupModel.Slogan, startupModel.Foundation)
		startupEntities = append(startupEntities, startup)
	}

	return startupEntities
}
