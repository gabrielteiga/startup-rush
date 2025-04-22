package requests

type BattleRequest struct {
	Battle []EventsRequest `json:"battle" validate:"required,min=2,dive"`
}

type EventsRequest struct {
	StartupID uint   `json:"startupId" validate:"required"`
	EventIDs  []uint `json:"eventIds" validate:"required"`
}

func NewBattleRequest(events []EventsRequest) *BattleRequest {
	return &BattleRequest{
		Battle: events,
	}
}
