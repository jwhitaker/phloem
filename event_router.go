package phloem

import "log"

type eventMappings map[EventIdentifier]EventHandler

type EventRouter struct {
	mappings eventMappings
}

func NewEventRouter() EventRouter {
	return EventRouter{ eventMappings {} }
}

func (eventRouter EventRouter) AddHandler(event string, aggregate string, eventHandler EventHandler) {
	log.Printf("Registering handler %s-%s-%s", event, aggregate, eventHandler)

	key := EventIdentifier{event, aggregate }

	eventRouter.mappings[key] = eventHandler
}

func (eventRouter EventRouter) Events() []EventIdentifier {
	eventKeys := make([]EventIdentifier, 0, len(eventRouter.mappings))

	for k := range eventRouter.mappings {
		eventKeys = append(eventKeys, k)
	}

	return eventKeys
}