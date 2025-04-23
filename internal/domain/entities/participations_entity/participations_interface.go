package participations_entity

type IParticipationRepository interface {
	Create(tournamentID, startupID uint, score int) (*Participation, error)
	FindByID(id uint) (*Participation, error)
	FindByTournamentID(tournamentID uint) ([]*Participation, error)
	FindByStartupID(startupID uint) ([]*Participation, error)
	AddScore(tournamentID, startupID uint, score int) error
}
