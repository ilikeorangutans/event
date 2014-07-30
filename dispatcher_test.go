package event

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDispatcherSubscribeShouldAddListener(t *testing.T) {
	callbackCalled := false

	ed := NewEventDispatcher()
	ed.Subscribe("foo", &FuncListener{Callback: func(e Event) bool { callbackCalled = true; return false }})

	ed.Announce(NewEvent("foo", nil))

	assert.True(t, callbackCalled, "Callback should have been called")
}
