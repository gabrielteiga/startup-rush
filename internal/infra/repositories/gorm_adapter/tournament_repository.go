package gorm_adapter

import (
	"log"

	"github.com/gabrielteiga/startup-rush/database"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/participations_entity"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/startup_entity"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/tournament_entity"
	"gorm.io/gorm"
)

type TournamentGORMRepository struct {
	DB *gorm.DB
}

func NewTournamentGORMRepository(db *gorm.DB) *TournamentGORMRepository {
	return &TournamentGORMRepository{
		DB: db,
	}
}

func (sr *TournamentGORMRepository) Create(startups []*startup_entity.Startup) *tournament_entity.Tournament {
	var tournament *database.Tournament
	var participantsEntity []*participations_entity.Participation

	err := sr.DB.Transaction(func(transaction *gorm.DB) error {
		tournament = &database.Tournament{
			Finished: false,
		}

		if err := transaction.Create(&tournament).Error; err != nil {
			log.Println("Error creating tournament: ", err)
			return err
		}

		for _, startup := range startups {
			participant := &database.StartupsTournaments{
				StartupID:    startup.ID,
				TournamentID: tournament.ID,
				Score:        70,
			}

			if err := transaction.Create(&participant).Error; err != nil {
				log.Println("Error creating participation: ", err)
				return err
			}

			participantsEntity = append(participantsEntity, participations_entity.NewParticipation(participant.ID, participant.Score, participant.StartupID, participant.TournamentID))
		}
		return nil
	})

	if err != nil {
		log.Println("Error creating tournament: ", err)
		return nil
	}

	return tournament_entity.NewTournament(tournament.ID, tournament.Finished, tournament.ChampionID, participantsEntity)
}
