package phloem

import "encoding/json"

// Event the event
type Event struct {
	Event   string
	Payload []byte
}

func NewEvent(event string, payload interface{}) (*Event, error) {
	rawPayload, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	return &Event {
		Event: event,
		Payload: rawPayload,
	}, nil
}

func (event Event) GetPayload(obj interface{}) error {
	return json.Unmarshal([]byte(event.Payload), obj)
}
