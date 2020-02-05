package phloem

import (
	 "sync"
	 "log"
)

// EventBus structure to hold event bus stuff
type EventBus struct {
	subscribers map[string]EventHandlerSlice
	mutex       sync.RWMutex
}

// NewEventBus creates a new event bus
func NewEventBus() EventBus {
	return EventBus{
		subscribers: map[string]EventHandlerSlice{},
	}
}

// Subscribe will call the handler when 
func (eventBus *EventBus) Subscribe(topic string, handler EventHandler) {
	eventBus.mutex.Lock()

	if prev, found := eventBus.subscribers[topic]; found {
		eventBus.subscribers[topic] = append(prev, handler)
	} else {
		eventBus.subscribers[topic] = append([]EventHandler{}, handler)
	}

	eventBus.mutex.Unlock()
}

// Publish will publish the specified message to the topic
func (eventBus *EventBus) Publish(topic string, event *Event) {
	eventBus.mutex.RLock()

	log.Printf("Publishing event %s to %s", event, topic)

	if _handlers, found := eventBus.subscribers[topic]; found {
		handlers := append(EventHandlerSlice{}, _handlers...)

		go func(msg *Event, handlerSlices EventHandlerSlice) {
			for _, h := range handlerSlices {
				h(msg)
			}
		}(event, handlers)
	}

	eventBus.mutex.RUnlock()
}
