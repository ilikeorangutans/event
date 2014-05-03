package observer

import (
	"github.com/ilikeorangutans/event"
)

// An entity that emits events to subscribed listeners
type SubscriberUnsubscriber interface {
	Subscribe(event.Listener)
	Unsubscribe(event.Listener)
}

type Announcer interface {
	Announce(event.Event)
}

// Helper type that takes care of the logistics of subscriber lists.
type Observable interface {
	Announcer
	SubscriberUnsubscriber
}

// Return a new support instance
func NewObservable() Observable {
	return &observableSupportImpl{}
}

type observableSupportImpl struct {
	subscribers []event.Listener
}

func (os *observableSupportImpl) Subscribe(listener event.Listener) {
	os.subscribers = append(os.subscribers, listener)
}

func (os *observableSupportImpl) Unsubscribe(listener event.Listener) {
	for i := range os.subscribers {
		subscriber := os.subscribers[i]

		if subscriber == listener {
			os.subscribers = append(os.subscribers[:i], os.subscribers[i+1:]...)
			return
		}
	}
}

func (os *observableSupportImpl) Announce(e event.Event) {
	for i := range os.subscribers {
		listener := os.subscribers[i]
		stopPropagation := listener.OnEvent(e)
		if stopPropagation {
			return
		}
	}
}
