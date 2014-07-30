package event

import (
	"fmt"
)

// Event type.
type Event interface {
	// Returns the source of this event. Might return nil.
	Source() interface{}
	// Returns a string identifying the event type.
	Type() string
}

// Creates a new event instance.
func NewEvent(t string, source interface{}) Event {
	return &eventImpl{
		source:    source,
		eventType: t,
	}
}

// Event implementation.
type eventImpl struct {
	eventType string
	source    interface{}
}

func (e *eventImpl) Type() string {
	return e.eventType
}

func (e *eventImpl) Source() interface{} {
	return e.source
}

func (e *eventImpl) String() string {
	return fmt.Sprintf("EventImpl{type=%s}", e.eventType)
}

// Helper listener that will log all events.
type LoggingListener struct{}

func (ll *LoggingListener) OnEvent(e Event) bool {
	fmt.Println(e)
	return false
}
