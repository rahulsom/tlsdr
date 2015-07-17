package main
import "container/list"

type Connection struct {
	success  bool
	events   list.List
	srcHost  string
	destHost string
}

type Event struct {
	success     bool
	eventType   string
	c2s         bool
	code 	    int
}

func NewConnection(from string, to string) Connection {
	return Connection{success:true, events:list.List{}, srcHost:from, destHost:to}
}

func NewEvent(code int) Event {
	// TODO Fix these things
	lookup := make(map[int]string)
	lookup[3] = "Client Hello"

	if code == 3 || code == 4 || code == 7 {
		return Event{success:true, eventType:lookup[code], c2s: true, code: code}
	} else {
		return Event{success:true, eventType:lookup[code], c2s: false, code: code}
	}
}

func AddEvent(connection Connection, event Event) {
	connection.events.PushBack(event)
}