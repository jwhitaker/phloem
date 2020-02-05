package phloem

// MemoryProducer a memory producer
type MemoryProducer struct {
	eventBus *EventBus
}

// NewMemoryProducer create a new memory producer
func NewMemoryProducer(eventBus *EventBus) MemoryProducer {
	return MemoryProducer{eventBus}
}

func (memoryProducer MemoryProducer) Send(topic string, event *Event) {
	memoryProducer.eventBus.Publish(topic, event)
}
