package gorm_adapter

import (
	"log"

	"github.com/gabrielteiga/startup-rush/database"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/battle_entity"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/participations_entity"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/startup_entity"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/tournament_entity"
	"gorm.io/gorm"
)

type TournamentGORMRepository struct {
	DB                       *gorm.DB
	IParticipationRepository participations_entity.IParticipationRepository
	IBattleRepository        battle_entity.IBattleRepository
}

func NewTournamentGORMRepository(db *gorm.DB, participationRepository participations_entity.IParticipationRepository, battleRepository battle_entity.IBattleRepository) *TournamentGORMRepository {
	return &TournamentGORMRepository{
		DB:                       db,
		IParticipationRepository: participationRepository,
		IBattleRepository:        battleRepository,
	}
}

func (sr *TournamentGORMRepository) Create(startups []*startup_entity.Startup) (*tournament_entity.Tournament, error) {
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

			participantsEntity = append(participantsEntity, participations_entity.NewParticipation(participant.ID, participant.StartupID, participant.TournamentID, participant.Score))
		}
		return nil
	})

	if err != nil {
		log.Println("Error creating tournament: ", err)
		return nil, err
	}

	return tournament_entity.NewTournament(tournament.ID, tournament.Finished, tournament.ChampionID, participantsEntity, nil), nil
}

func (sr *TournamentGORMRepository) List() ([]*tournament_entity.Tournament, error) {
	var tournaments []database.Tournament
	if err := sr.DB.
		Preload("Startups").
		Preload("Battles").
		Preload("Battles.Startup1").
		Preload("Battles.Startup2").
		Preload("Battles.BattleEvents.Event").
		Find(&tournaments).
		Error; err != nil {
		return nil, err
	}

	var tournamentsEntity []*tournament_entity.Tournament
	for i := range tournaments {
		participants, err := sr.IParticipationRepository.FindByTournamentID(tournaments[i].ID)
		if err != nil {
			log.Println("Error finding participants: ", err)
			return nil, err
		}

		battles, err := sr.IBattleRepository.FindByTournamentID(tournaments[i].ID)
		if err != nil {
			log.Println("Error finding battles: ", err)
			return nil, err
		}

		tournamentEntity := tournament_entity.NewTournament(
			tournaments[i].ID,
			tournaments[i].Finished,
			tournaments[i].ChampionID,
			participants,
			battles,
		)

		tournamentsEntity = append(tournamentsEntity, tournamentEntity)
	}
	return tournamentsEntity, nil
}

func (sr *TournamentGORMRepository) FindByID(id uint) (*tournament_entity.Tournament, error) {
	var tournament *database.Tournament

	if err := sr.DB.Where("id = ?", id).First(&tournament).Error; err != nil {
		log.Println("Tournament not found")
		return nil, err
	}

	participantsEntity, err := sr.IParticipationRepository.FindByTournamentID(id)
	if err != nil {
		log.Println("Error finding participants: ", err)
		return nil, err
	}

	battles, err := sr.IBattleRepository.FindByTournamentID(id)
	if err != nil {
		log.Println("Error finding battles: ", err)
		return nil, err
	}

	return tournament_entity.NewTournament(tournament.ID, tournament.Finished, tournament.ChampionID, participantsEntity, battles), nil
}
