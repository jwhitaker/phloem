package phloem

import "log"

type eventMappings map[string]EventHandler

type EventRouter struct {
	mappings eventMappings
}

func NewEventRouter() EventRouter {
	return EventRouter{ eventMappings {} }
}

func (eventRouter EventRouter) AddHandler(event string, eventHandler EventHandler) {
	log.Printf("Registering handler %s", event, eventHandler)

	eventRouter.mappings[event] = eventHandler
}

func (eventRouter EventRouter) Events() []string {
	eventKeys := make([]string, 0, len(eventRouter.mappings))

	for k := range eventRouter.mappings {
		eventKeys = append(eventKeys, k)
	}

	return eventKeys
}

func (eventRouter EventRouter) GetHandler(event string) (EventHandler, bool) {
	handler, ok := eventRouter.mappings[event]

	return handler, ok
}
