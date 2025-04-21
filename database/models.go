package database

import (
	"time"

	"gorm.io/gorm"
)

type Startup struct {
	gorm.Model
	Name       string    `gorm:"not null"`
	Slogan     string    `gorm:"not null"`
	Foundation time.Time `gorm:"not null"`

	Tournaments []Tournament `gorm:"many2many:startups_tournaments"`
}

type Tournament struct {
	gorm.Model
	Finished bool `gorm:"not null;default:false"`

	ChampionID *uint
	Champion   *Startup  `gorm:"foreignKey:ChampionID;constraint;OnDelete:SET NULL"`
	Startups   []Startup `gorm:"many2many:startups_tournaments"`
}

type StartupsTournaments struct {
	gorm.Model
	StartupID    uint `gorm:"not null"`
	TournamentID uint `gorm:"not null"`
	Score        uint `gorm:"not null;default:0"`

	Startup    *Startup    `gorm:"foreignKey:StartupID"`
	Tournament *Tournament `gorm:"foreignKey:TournamentID"`
}

type Battle struct {
	gorm.Model
	TournamentID      uint `gorm:"not null"`
	Startup1ID        uint `gorm:"not null"`
	Startup2ID        uint `gorm:"not null"`
	Finished          bool `gorm:"not null;default:false"`
	Startup1Score     *uint
	Startup2Score     *uint
	BattleChildren1ID *uint
	BattleChildren2ID *uint
	WinnerID          *uint

	Tournament      *Tournament `gorm:"foreignKey:TournamentID"`
	Startup1        *Startup    `gorm:"foreignKey:Startup1ID"`
	Startup2        *Startup    `gorm:"foreignKey:Startup2ID"`
	BattleChildren1 *Battle     `gorm:"foreignKey:BattleChildren1ID;constraint:OnDelete:SET NULL"`
	BattleChildren2 *Battle     `gorm:"foreignKey:BattleChildren2ID;constraint:OnDelete:SET NULL"`
	Winner          *Startup    `gorm:"foreignKey:WinnerID;constraint:OnDelete:SET NULL"`
}

type Events struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Score uint   `gorm:"not null"`
}

type BattlesEvents struct {
	gorm.Model
	BattleID  uint `gorm:"not null"`
	EventID   uint `gorm:"not null"`
	StartupID uint `gorm:"not null"`
	Checked   bool `gorm:"not null;default:false"`
}
