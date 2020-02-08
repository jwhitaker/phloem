package phloem

import "encoding/json"

// Event the event
type Event struct {
	Event   string
	Payload []byte
}

// NewEvent creates a new event
func NewEvent(event string, payload interface{}) (*Event, error) {
	rawPayload, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	return &Event{
		Event:   event,
		Payload: rawPayload,
	}, nil
}

// GetPayload gets the payload from an event
func (event Event) GetPayload(obj interface{}) error {
	return json.Unmarshal([]byte(event.Payload), obj)
}
