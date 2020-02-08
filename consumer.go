package phloem

// EventHandler defines a handler
type EventHandler func(event *Event)

// EventHandlerSlice defines a slice of handlers
type EventHandlerSlice []EventHandler

// Consumer defines a consumer
type Consumer interface {
	Subscribe(events []string)
	Poll() *Event
	Close()
}
