package phloem

// Producer defines an event publisher
type Producer interface {
	Send(event Event)
}
