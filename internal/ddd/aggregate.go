package ddd

type Aggregate interface {
	GetId() string
	AddEvent(event Event)
	GetEvents() []Event
}

type BaseAggregate struct {
	Id     string
	Events []Event
}

func (aggregate BaseAggregate) GetId() string {
	return aggregate.Id
}

func (aggregate BaseAggregate) GetEvents() []Event {
	return aggregate.Events
}

func (aggregate *BaseAggregate) AddEvent(event Event) {
	aggregate.Events = append(aggregate.Events, event)
}
