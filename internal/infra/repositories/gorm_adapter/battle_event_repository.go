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
