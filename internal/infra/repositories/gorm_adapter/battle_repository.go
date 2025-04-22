package gorm_adapter

import (
	"github.com/gabrielteiga/startup-rush/database"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/battle_entity"
	"gorm.io/gorm"
)

type BattleGORMRepository struct {
	DB *gorm.DB
}

func NewBattleGORMRepository(db *gorm.DB) *BattleGORMRepository {
	return &BattleGORMRepository{
		DB: db,
	}
}

func (br *BattleGORMRepository) Create(tournamentID, startup1ID, startup2ID uint, battleChildren1ID, battleChildren2ID *uint, phase battle_entity.BattlePhase) (*battle_entity.Battle, error) {
	battle := &database.Battle{
		TournamentID:      tournamentID,
		Startup1ID:        startup1ID,
		Startup2ID:        startup2ID,
		BattleChildren1ID: battleChildren1ID,
		BattleChildren2ID: battleChildren2ID,
		Phase:             phase,
	}

	if err := br.DB.Create(&battle).Error; err != nil {
		return nil, err
	}

	return battle_entity.NewBattle(
		battle.ID,
		battle.TournamentID,
		battle.Startup1ID,
		battle.Startup2ID,
		nil,
		nil,
		battle.Finished,
		nil,
		battle.BattleParentID,
		battle.BattleChildren1ID,
		battle.BattleChildren2ID,
		battle.Phase,
	), nil
}

func (br *BattleGORMRepository) FindByID(id uint) (*battle_entity.Battle, error) {
	battle := &database.Battle{}
	if err := br.DB.Preload("Startup1").Preload("Startup2").Preload("BattleParent").Preload("BattleChildren1").Preload("BattleChildren2").First(&battle, id).Error; err != nil {
		return nil, err
	}

	return battle_entity.NewBattle(
		battle.ID,
		battle.TournamentID,
		battle.Startup1ID,
		battle.Startup2ID,
		battle.ScoreStartup1,
		battle.ScoreStartup2,
		battle.Finished,
		battle.WinnerID,
		battle.BattleParentID,
		battle.BattleChildren1ID,
		battle.BattleChildren2ID,
		battle.Phase,
	), nil
}

func (br *BattleGORMRepository) FindByTournamentID(tournamentID uint) ([]*battle_entity.Battle, error) {
	var battles []*database.Battle
	if err := br.DB.Where("tournament_id = ?", tournamentID).Find(&battles).Error; err != nil {
		return nil, err
	}

	var battlesEntity []*battle_entity.Battle
	for _, battle := range battles {
		battlesEntity = append(battlesEntity, battle_entity.NewBattle(
			battle.ID,
			battle.TournamentID,
			battle.Startup1ID,
			battle.Startup2ID,
			battle.ScoreStartup1,
			battle.ScoreStartup2,
			battle.Finished,
			battle.WinnerID,
			battle.BattleParentID,
			battle.BattleChildren1ID,
			battle.BattleChildren2ID,
			battle.Phase,
		))
	}

	return battlesEntity, nil
}
