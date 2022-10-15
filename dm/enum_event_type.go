package dm

type EventType string

const (
	EventTypeActual    EventType = "Actual"
	EventTypeExpected  EventType = "Expected"
	EventTypePlanned   EventType = "Planned"
	EventTypeRequested EventType = "Requested"
)
