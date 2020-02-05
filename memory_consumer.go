package phloem

type MemoryConsumer struct {
	eventBus *EventBus
}

func NewMemoryConsumer(eventBus *EventBus) MemoryConsumer {
	return MemoryConsumer{ eventBus }
}

func (memoryConsumer MemoryConsumer) Subscribe(topic string, handler EventHandler) {
	memoryConsumer.eventBus.Subscribe(topic, handler)
}
