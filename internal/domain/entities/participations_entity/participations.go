package participations_entity

type Participation struct {
	ID           uint `json:"id,omitempty"`
	StartupID    uint `json:"startupId"`
	TournamentID uint `json:"tournamentId"`
	Score        int  `json:"score"`
}

func NewParticipation(id, startupId, tournamentId uint, score int) *Participation {
	return &Participation{
		ID:           id,
		StartupID:    startupId,
		TournamentID: tournamentId,
		Score:        score,
	}
}
