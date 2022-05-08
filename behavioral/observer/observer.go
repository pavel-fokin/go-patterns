package main

// Event defines an indication of a some occurence
type Event struct {
	// Data in this case is a simple int.
	Data int
}


// Observer defines a standard interface to listen for a specific event.
type Observer interface {
	// OnNotify allows to publsh an event
	OnNotify(Event)
}

// Notifier is the instance being observed.
type  Notifier interface {
	// Register itself to listen/observe events.
	Register(Observer)
	// Remove itself from the collection of observers/listeners.
	Unregister(Observer)
	// Notify publishes new events to listeners.
	Notify(Event)
}
