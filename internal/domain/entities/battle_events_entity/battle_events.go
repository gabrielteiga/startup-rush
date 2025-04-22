package battle_events_entity

type BattleEvents struct {
	ID              uint `json:"id,omitempty"`
	ParticipationID uint `json:"participationId"`
	BattleID        uint `json:"battleId"`
	EventID         uint `json:"eventId"`
	Checked         bool `json:"checked"`
}
