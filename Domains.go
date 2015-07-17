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

func NewEvent(ucode uint8) Event {
	code := int(ucode)
	// TODO Fix these things
	lookup := make(map[int]string)

	lookup[20] = "TypeChangeCypherSpec"
	lookup[21] = "TypeAlert"
	lookup[22] = "TypeHandshake"
	lookup[23] = "TypeApplicationData"
	lookup[0] = "HandshakeTypeHelloRequest"
	lookup[1] = "HandshakeTypeClientHello"
	lookup[2] = "HandshakeTypeServerHello"
	lookup[3] = "HandshakeTypeHelloVerifyRequest"
	lookup[11] = "HandshakeTypeCertificate"
	lookup[12] = "HandshakeTypeServerKeyExchange"
	lookup[13] = "HandshakeTypeCertificateRequest"
	lookup[14] = "HandshakeTypeServerHelloDone"
	lookup[15] = "HandshakeTypeCertificateVerify"
	lookup[16] = "HandshakeTypeClientKeyExchange"
	lookup[20] = "HandshakeTypeFinished"

	// TODO FIx this
	if code == 22 || code == 11 || code == 13 {
		return Event{success:true, eventType:lookup[code], c2s: true, code: code}
	} else {
		return Event{success:true, eventType:lookup[code], c2s: false, code: code}
	}
}

func (event Event) String() string {
	return fmt.Sprintf("Event{success: %t, type: '%s'($d)}", event.success, event.eventType, event.code)
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