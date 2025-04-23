package gorm_adapter

import (
	"github.com/gabrielteiga/startup-rush/database"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/battle_events_entity"
	"gorm.io/gorm"
)

type BattleEventGORMRepository struct {
	DB *gorm.DB
}

func NewBattleEventGORMRepository(db *gorm.DB) *BattleEventGORMRepository {
	return &BattleEventGORMRepository{
		DB: db,
	}
}

func (ber *BattleEventGORMRepository) Create(battleID, startupID, eventID uint) (*battle_events_entity.BattleEvents, error) {
	var battleEvent *database.BattlesEvents
	err := ber.DB.Where("battle_id = ? AND startup_id = ? AND event_id = ?", battleID, startupID, eventID).First(&battleEvent).Error
	if err != nil {
		battleEvent = &database.BattlesEvents{
			BattleID:  battleID,
			StartupID: startupID,
			EventID:   eventID,
			Checked:   true,
		}

		if err := ber.DB.Create(&battleEvent).Error; err != nil {
			return nil, err
		}
	}

	return battle_events_entity.NewBattleEvents(
		battleEvent.ID,
		battleEvent.StartupID,
		battleEvent.BattleID,
		battleEvent.EventID,
		battleEvent.Checked,
	), nil
}

func (ber *BattleEventGORMRepository) GetBattleDatabaseWithEvents(battleID uint) (*database.Battle, error) {
	var battle *database.Battle
	if err := ber.DB.Preload("BattleEvents").Preload("BattleEvents.Event").First(&battle, battleID).Error; err != nil {
		return nil, err
	}
	return battle, nil
}

func (ber *BattleEventGORMRepository) CountEventsByTournament(tournamentID uint) ([]*battle_events_entity.EventStat, error) {
	var events []*battle_events_entity.EventStat
	if err := ber.DB.
		Table("battles_events be").
		Select("be.startup_id AS startup_id, e.name AS event_name, COUNT(*) AS total").
		Joins("JOIN events e ON e.id = be.event_id").
		Joins("JOIN battles b ON b.id = be.battle_id").
		Where("b.tournament_id = ? AND be.checked = ?", tournamentID, true).
		Group("be.startup_id, e.name").
		Scan(&events).
		Error; err != nil {
		return nil, err
	}

	return events, nil
}
