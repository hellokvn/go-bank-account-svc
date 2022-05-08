package cqrs

type BaseEvent struct {
	Id string
}

func (e *BaseEvent) GetEventId() string {
	return e.Id
}

func (e *BaseEvent) SetEventId(id string) {
	e.Id = id
}
