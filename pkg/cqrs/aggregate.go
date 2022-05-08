package cqrs

const (
	aggregateDefaultVersion = -1
)

type AggregateType string

// type AggregateRoot interface {
// 	GetUncommittedEvents() []Event
// 	GetID() string
// 	SetID(id string) *AggregateBase
// 	GetVersion() int64
// 	ClearUncommittedEvents()
// 	ToSnapshot()
// 	SetType(aggregateType AggregateType)
// 	GetType() AggregateType
// 	SetAppliedEvents(events []Event)
// 	GetAppliedEvents() []Event
// 	RaiseEvent(event Event) error
// 	String() string
// 	Load
// 	Apply
// }

type BaseAggregate struct {
	Id      string
	Version int32
	Type    AggregateType
	Changes []interface{}
}

// Getter & Setter

func (a *BaseAggregate) GetAggregateId() string {
	return a.Id
}

func (a *BaseAggregate) SetAggregateId(id string) {
	a.Id = id
}

func (a *BaseAggregate) GetAggregateType() AggregateType {
	return a.Type
}

func (a *BaseAggregate) SetAggregateType(aggregateType AggregateType) {
	a.Type = aggregateType
}

func (a *BaseAggregate) GetAggregateVersion() int32 {
	return a.Version
}

func (a *BaseAggregate) SetAggregateVersion(version int32) {
	a.Version = version
}

func (a *BaseAggregate) ApplyChange(event interface{}, isNewEvent bool) error {
	if isNewEvent {
		a.Changes = append(a.Changes, event)
	}

	return nil
}
