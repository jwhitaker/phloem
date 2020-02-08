package phloem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var eventName1 = "EventName1"
var eventName2 = "EventName2"
var noOp = func(_ *Event) {}

func createRouter() EventRouter {
	router := NewEventRouter()
	router.AddHandler(eventName2, noOp)
	router.AddHandler(eventName1, noOp)

	return router
}

func TestEventsNotEmpty(t *testing.T) {
	router := createRouter()

	events := router.Events()

	expectedEvents := []string{eventName2, eventName1}

	assert.Equal(t, expectedEvents, events)
}

func TestEventsEmpty(t *testing.T) {
	router := NewEventRouter()

	events := router.Events()

	assert.Equal(t, []string{}, events)
}

func TestGetHandlerExists(t *testing.T) {
	router := createRouter()

	handler, ok := router.GetHandler(eventName1)
	assert.True(t, ok, "Handler does not exist")
	assert.NotNil(t, handler, "Handler is nil")
}

func TestGetHandlerNotExists(t *testing.T) {
	router := createRouter()

	handler, ok := router.GetHandler("UnknownHandler")

	assert.False(t, ok)
	assert.Nil(t, handler)
}
