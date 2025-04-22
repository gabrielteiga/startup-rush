package event_entity

type IEventRepository interface {
	Create(name string, score int) *Event
	List() []*Event
}
