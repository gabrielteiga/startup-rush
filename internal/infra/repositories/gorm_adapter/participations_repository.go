package gorm_adapter

import (
	"log"

	"github.com/gabrielteiga/startup-rush/database"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/battle_entity"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/participations_entity"
	"gorm.io/gorm"
)

type ParticipationsGORMRepository struct {
	DB                       *gorm.DB
	IParticipationRepository participations_entity.IParticipationRepository
	IBattleRepository        battle_entity.IBattleRepository
}

func NewParticipationsGORMRepository(db *gorm.DB) *ParticipationsGORMRepository {
	return &ParticipationsGORMRepository{
		DB: db,
	}
}

func (pr *ParticipationsGORMRepository) Create(tournamentID, startupID uint, score int) (*participations_entity.Participation, error) {
	participation := &database.StartupsTournaments{
		TournamentID: tournamentID,
		StartupID:    startupID,
		Score:        score,
	}

	if err := pr.DB.Create(&participation).Error; err != nil {
		return nil, err
	}

	return participations_entity.NewParticipation(participation.ID, participation.StartupID, participation.TournamentID, participation.Score), nil
}

func (pr *ParticipationsGORMRepository) FindByID(id uint) (*participations_entity.Participation, error) {
	var participation database.StartupsTournaments

	if err := pr.DB.First(&participation, id).Error; err != nil {
		return nil, err
	}

	return participations_entity.NewParticipation(participation.ID, participation.StartupID, participation.TournamentID, participation.Score), nil
}

func (pr *ParticipationsGORMRepository) FindByTournamentID(tournamentID uint) ([]*participations_entity.Participation, error) {
	var participations []database.StartupsTournaments

	if err := pr.DB.Where("tournament_id = ?", tournamentID).Find(&participations).Error; err != nil {
		return nil, err
	}

	var participationsEntity []*participations_entity.Participation
	for _, participation := range participations {
		participationsEntity = append(participationsEntity, participations_entity.NewParticipation(participation.ID, participation.StartupID, participation.TournamentID, participation.Score))
	}

	return participationsEntity, nil
}

func (pr *ParticipationsGORMRepository) FindByStartupID(startupID uint) ([]*participations_entity.Participation, error) {
	var participations []database.StartupsTournaments

	if err := pr.DB.Where("startup_id = ?", startupID).Find(&participations).Error; err != nil {
		return nil, err
	}

	var participationsEntity []*participations_entity.Participation
	for _, participation := range participations {
		participationsEntity = append(participationsEntity, participations_entity.NewParticipation(participation.ID, participation.StartupID, participation.TournamentID, participation.Score))
	}

	return participationsEntity, nil
}

func (pr *ParticipationsGORMRepository) AddScore(tournamentID, startupID uint, score int) error {
	var participation database.StartupsTournaments
	if err := pr.DB.Where("tournament_id = ? AND startup_id = ?", tournamentID, startupID).First(&participation).Error; err != nil {
		log.Println("Error finding participation:", err)
		return err
	}

	participation.Score += score

	return pr.DB.Model(&database.StartupsTournaments{}).Select("score").Updates(participation).Error
}
