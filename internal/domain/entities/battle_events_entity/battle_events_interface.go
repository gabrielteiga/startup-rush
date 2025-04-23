package battle_events_entity

import "github.com/gabrielteiga/startup-rush/database"

type EventStat struct {
	StartupID uint
	EventName string
	Total     int
}

type IBattleEventsRepository interface {
	Create(battleID, startupID, eventID uint) (*BattleEvents, error)
	GetBattleDatabaseWithEvents(battleID uint) (*database.Battle, error)
	CountEventsByTournament(tournamentID uint) ([]*EventStat, error)
}
