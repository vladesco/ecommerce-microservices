package ddd

type Aggregate interface {
	GetId() string
	AddEvent(event Event)
	GetEvents() []Event
}

type BaseAggregate struct {
	id     string
	events []Event
}

func (aggregate BaseAggregate) GetId() string {
	return aggregate.id
}

func (aggregate BaseAggregate) GetEvents() []Event {
	return aggregate.events
}

func (aggregate *BaseAggregate) AddEvent(event Event) {
	aggregate.events = append(aggregate.events, event)
}
