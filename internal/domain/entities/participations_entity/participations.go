package participations_entity

type Participation struct {
	ID           uint `json:"id,omitempty"`
	StartupID    uint `json:"startupId"`
	TournamentID uint `json:"tournamentId"`
	Score        uint `json:"score"`
}

func NewParticipation(id, score, startupId, tournamentId uint) *Participation {
	return &Participation{
		ID:           id,
		StartupID:    startupId,
		TournamentID: tournamentId,
		Score:        score,
	}
}
