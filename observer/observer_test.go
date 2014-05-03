package observer

import (
	"github.com/ilikeorangutans/event"
	"testing"
)

func TestFoo(t *testing.T) {

	oh := NewObservable()

	oh.Announce(event.NewEvent("type", nil))
	println(oh)
}

type TestListener struct {
	numEvents int
}

func (tl *TestListener) OnEvent(e event.Event) bool {
	tl.numEvents++
	return false
}

func TestSubscribe(t *testing.T) {

	tl := &TestListener{}
	oh := NewObservable()

	oh.Subscribe(tl)
	oh.Announce(event.NewEvent("type", nil))

	if tl.numEvents != 1 {
		t.Error("We should have received one event")
	}

}
