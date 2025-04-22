package battle_events_entity

type BattleEvents struct {
	ID              uint `json:"id,omitempty"`
	ParticipationID uint `json:"participationId"`
	BattleID        uint `json:"battleId"`
	EventID         uint `json:"eventId"`
	Checked         bool `json:"checked"`
}

func NewBattleEvents(id, participationID, battleID, eventID uint, checked bool) *BattleEvents {
	return &BattleEvents{
		ID:              id,
		ParticipationID: participationID,
		BattleID:        battleID,
		EventID:         eventID,
		Checked:         checked,
	}
}
