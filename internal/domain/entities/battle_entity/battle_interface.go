package battle_entity

type IBattleRepository interface {
	Create(tournamentID, startup1ID, startup2ID uint, battleChildren1ID, battleChildren2ID *uint, phase BattlePhase) (*Battle, error)
	FindByID(id uint) (*Battle, error)
	FindByTournamentID(tournamentID uint) ([]*Battle, error)
}
