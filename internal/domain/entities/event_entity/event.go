package event_entity

type Event struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func NewEvent(id uint, name string, score int) *Event {
	return &Event{
		ID:    id,
		Name:  name,
		Score: score,
	}
}
