# event

Really simple event library written in go. 

## Usage 


	package main
	
	import (
		"github.com/ilikeorangutans/event"
		"github.com/ilikeorangutans/event/observer"
	)
	
	func main() {
		observable := NewObservable()
		listener := &LoggingListener{}
		
		observable.Subscribe(listener)
		observable.Announce(event.NewEvent("my event type", nil))
	}
	
Output: 
	
	EventImpl{type=my event type}