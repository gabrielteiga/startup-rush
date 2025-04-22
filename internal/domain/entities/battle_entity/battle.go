package battle_entity

type Battle struct {
	ID            uint  `json:"id"`
	TournamentID  uint  `json:"tournamentId"`
	Startup1ID    uint  `json:"startup1Id"`
	Startup2ID    uint  `json:"startup2Id"`
	ScoreStartup1 *int  `json:"score1"`
	ScoreStartup2 *int  `json:"score2"`
	Finished      bool  `json:"finished"`
	WinnerID      *uint `json:"winnerId"`

	BattleParentID    *uint `json:"battleParentId,omitempty"`
	BattleChildren1ID *uint `json:"battleChildren1Id,omitempty"`
	BattleChildren2ID *uint `json:"battleChildren2Id,omitempty"`

	Phase           BattlePhase `json:"phase"`
	ParentBattle    *Battle     `json:"parentBattle,omitempty"`
	ChildrenBattles []*Battle   `json:"childrenBattles,omitempty"`
}

type BattlePhase string

const (
	PhaseQuarterFinal BattlePhase = "quarter_final"
	PhaseSemiFinal    BattlePhase = "semi_final"
	PhaseFinal        BattlePhase = "final"
)

func NewBattle(id, tournamentID, startup1ID, startup2ID uint, score1, score2 *int, finished bool, winnerId, battleParentID, battleChildren1ID, battleChildren2ID *uint, phase BattlePhase) *Battle {
	return &Battle{
		ID:                id,
		TournamentID:      tournamentID,
		Startup1ID:        startup1ID,
		Startup2ID:        startup2ID,
		WinnerID:          winnerId,
		ScoreStartup1:     score1,
		ScoreStartup2:     score2,
		Phase:             phase,
		Finished:          finished,
		BattleParentID:    battleParentID,
		BattleChildren1ID: battleChildren1ID,
		BattleChildren2ID: battleChildren2ID,
	}
}
