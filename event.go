package event

import (
	"fmt"
)

type Event interface {
	Source() interface{}
	Type() string
}

func NewEvent(t string, source interface{}) Event {
	return &eventImpl{
		source:    source,
		eventType: t,
	}
}

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

type Listener interface {
	// Callback called for events. Returns true if event handling should stop after this handler.
	OnEvent(Event) bool
}

type ListenerFunc func(Event) bool
