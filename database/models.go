package database

import (
	"time"

	"github.com/gabrielteiga/startup-rush/internal/domain/entities/battle_entity"
	"gorm.io/gorm"
)

type Startup struct {
	gorm.Model
	Name       string    `gorm:"not null"`
	Slogan     string    `gorm:"not null"`
	Foundation time.Time `gorm:"not null"`

	Participation []StartupsTournaments `gorm:"foreignKey:StartupID"`
}

type Tournament struct {
	gorm.Model
	Finished bool `gorm:"not null;default:false"`

	ChampionID *uint
	Champion   *Startup              `gorm:"foreignKey:ChampionID;constraint;OnDelete:SET NULL"`
	Battles    []Battle              `gorm:"foreignKey:TournamentID"`
	Startups   []StartupsTournaments `gorm:"foreignKey:TournamentID"`
}

type StartupsTournaments struct {
	gorm.Model
	StartupID    uint `gorm:"not null"`
	TournamentID uint `gorm:"not null"`
	Score        int  `gorm:"not null;default:70"`

	Startup    *Startup    `gorm:"foreignKey:StartupID"`
	Tournament *Tournament `gorm:"foreignKey:TournamentID"`
}

type Battle struct {
	gorm.Model
	TournamentID      uint `gorm:"not null"`
	Startup1ID        uint `gorm:"not null"`
	Startup2ID        uint `gorm:"not null"`
	Finished          bool `gorm:"not null;default:false"`
	ScoreStartup1     *int
	ScoreStartup2     *int
	BattleParentID    *uint
	BattleChildren1ID *uint
	BattleChildren2ID *uint
	WinnerID          *uint
	Phase             battle_entity.BattlePhase `gorm:"type:ENUM('quarter_final','semi_final','final');not null"`

	Tournament *Tournament `gorm:"foreignKey:TournamentID"`
	Startup1   *Startup    `gorm:"foreignKey:Startup1ID"`
	Startup2   *Startup    `gorm:"foreignKey:Startup2ID"`

	BattleParent    *Battle `gorm:"foreignKey:BattleParentID;constraint:OnDelete:SET NULL"`
	BattleChildren1 *Battle `gorm:"foreignKey:BattleChildren1ID;constraint:OnDelete:SET NULL"`
	BattleChildren2 *Battle `gorm:"foreignKey:BattleChildren2ID;constraint:OnDelete:SET NULL"`

	Winner *Startup `gorm:"foreignKey:WinnerID;constraint:OnDelete:SET NULL"`

	BattleEvents []BattlesEvents `gorm:"foreignKey:BattleID"`
}

type Events struct {
	gorm.Model
	Name  string `gorm:"size:255;not null;uniqueIndex"`
	Score int    `gorm:"not null"`

	Battles []BattlesEvents `gorm:"foreignKey:EventID"`
}

type BattlesEvents struct {
	gorm.Model
	BattleID  uint `gorm:"not null"`
	EventID   uint `gorm:"not null"`
	StartupID uint `gorm:"not null"`
	Checked   bool `gorm:"not null;default:false"`

	Battle  *Battle  `gorm:"foreignKey:BattleID"`
	Event   *Events  `gorm:"foreignKey:EventID"`
	Startup *Startup `gorm:"foreignKey:StartupID"`
}
