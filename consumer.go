package phloem

// EventIdentifier the event key
type EventIdentifier struct {
	Event string
	Aggregate string
}

// Event the event
type Event struct {
	Event   EventIdentifier
	Payload interface{}
}

// EventHandler defines a handler
type EventHandler func(event *Event)

// EventHandlerSlice defines a slice of handlers
type EventHandlerSlice []EventHandler

// Consumer defines a consumer
type Consumer interface {
	Subscribe(eventIds []EventIdentifier)
	Poll() *Event
	Close()
}
