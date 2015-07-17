package main
import (
	"container/list"
	"fmt"
)

type Connection struct {
	success         bool
	events          list.List
	srcHost         string
	destHost        string
	recommendations list.List
	failedReason    string
}

type Event struct {
	success   bool
	eventType string
	c2s       bool
	code      int
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

func (event Event) String() string {
	return fmt.Sprintf("Event{success: %t, type: '%s'}", event.success, event.eventType)
}

func (connection Connection) String() string {
	return fmt.Sprintf("Connection{success: %t, failReason: '%s', src: '%s', dest: '%s', recommendations: %#v, events: %#v}",
		connection.success, connection.failedReason, connection.srcHost, connection.destHost, connection.recommendations, connection.events)
}

func (connection Connection) AddEvent(event Event) {
	connection.events.PushBack(event)
}

func (connection Connection) WithEvent(event Event) Connection {
	connection.events.PushBack(event)
	return connection
}