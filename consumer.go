package phloem

// Event the event
type Event struct {
	Payload string
}

// EventHandler defines a handler
type EventHandler func(event *Event)

// EventHandlerSlice defines a slice of handlers
type EventHandlerSlice []EventHandler

// Consumer defines a consumer
type Consumer interface {
	Subscribe(topic string, handler EventHandler)
}
