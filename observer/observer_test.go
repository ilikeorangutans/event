package observer

import (
	"github.com/ilikeorangutans/event"
	"testing"
)

type TestListener struct {
	numEvents int
}

func (tl *TestListener) OnEvent(e event.Event) bool {
	tl.numEvents++
	return false
}

func TestSubscribe(t *testing.T) {
	oh := NewObservable()

	tl1 := &TestListener{}
	oh.Subscribe(tl1)
	tl2 := &TestListener{}
	oh.Subscribe(tl2)

	oh.Announce(event.NewEvent("type", nil))

	if tl1.numEvents != 1 || tl2.numEvents != 1 {
		t.Error("We should have received one event")
	}
}

func TestUnsubscribe(t *testing.T) {
	oh := NewObservable()

	tl1 := &TestListener{}
	oh.Subscribe(tl1)
	tl2 := &TestListener{}
	oh.Subscribe(tl2)

	oh.Unsubscribe(tl1)

	oh.Announce(event.NewEvent("type", nil))

	if tl1.numEvents != 0 || tl2.numEvents != 1 {
		t.Error("We should have received one event")
	}
}

func ExampleBasic() {

	observable := NewObservable()
	listener := &event.LoggingListener{}

	observable.Subscribe(listener)
	observable.Announce(event.NewEvent("my event type", nil))

	// Output: EventImpl{type=my event type}
}
