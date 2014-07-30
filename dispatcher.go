package event

// An entity that keeps track of event listeners
type SubscriberUnsubscriber interface {
	// Add the given listener to the subscribers for the given event type.
	Subscribe(string, Listener)
	// Unsubscribe the given listener.
	Unsubscribe(Listener)
}

type Announcer interface {
	Announce(Event)
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		subscribers: make(map[string][]Listener),
	}
}

type EventDispatcher struct {
	subscribers map[string][]Listener
}

func (ed *EventDispatcher) Announce(event Event) {
	subscribers, ok := ed.subscribers[event.Type()]

	if !ok {
		return
	}

	for i := range subscribers {
		subscribers[i].OnEvent(event)
	}
}

func (ed *EventDispatcher) Subscribe(eventType string, listener Listener) {
	_, ok := ed.subscribers[eventType]
	if !ok {
		ed.subscribers[eventType] = make([]Listener, 0)
	}

	ed.subscribers[eventType] = append(ed.subscribers[eventType], listener)
}
