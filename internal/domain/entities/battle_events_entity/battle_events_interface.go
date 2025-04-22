package battle_events_entity

import "github.com/gabrielteiga/startup-rush/database"

type IBattleEventsRepository interface {
	Create(battleID, startupID, eventID uint) (*BattleEvents, error)
	GetBattleDatabaseWithEvents(battleID uint) (*database.Battle, error)
}
