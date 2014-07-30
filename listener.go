package event

// Describes an event callback.
type Listener interface {
	// Callback called for events. Returns true if event handling should stop after this handler.
	OnEvent(Event) bool
}

// Simple wrapper for anonymous functions into a Listener.
type FuncListener struct {
	Callback func(Event) bool
}

func (fl *FuncListener) OnEvent(event Event) bool {
	return fl.Callback(event)
}
