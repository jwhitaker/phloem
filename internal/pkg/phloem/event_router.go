package phloem

import "log"

type eventMappings map[string]EventHandler

// EventRouter defines an event router
type EventRouter struct {
	mappings eventMappings
}

// NewEventRouter creates a new event router
func NewEventRouter() EventRouter {
	return EventRouter{eventMappings{}}
}

// AddHandler adds an event handler for an event
func (eventRouter EventRouter) AddHandler(event string, eventHandler EventHandler) {
	log.Printf("Registering handler for event %s", event)

	eventRouter.mappings[event] = eventHandler
}

// Events returns a list of handled events
func (eventRouter EventRouter) Events() []string {
	eventKeys := make([]string, 0, len(eventRouter.mappings))

	for k := range eventRouter.mappings {
		eventKeys = append(eventKeys, k)
	}

	return eventKeys
}

// GetHandler returns a handler for an event
func (eventRouter EventRouter) GetHandler(event string) (EventHandler, bool) {
	handler, ok := eventRouter.mappings[event]

	return handler, ok
}
