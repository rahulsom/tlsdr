package main
import (
	"container/list"
	"github.com/certifyTian/TLSHandshakeDecoder"
)

func createMutual(stages int) Connection {
	conn := NewConnection("localhost", "localhost:443")

	if (stages > 0) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeClientHello, true))}
	if (stages > 1) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeServerHello, false))}
	if (stages > 2) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeCertificate, false))}
	if (stages > 3) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeServerKeyExchange, false))}
	if (stages > 4) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeCertificateRequest, false))}
	if (stages > 5) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeServerHelloDone, false))}
	if (stages > 6) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeCertificate, true))}
	if (stages > 7) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeClientKeyExchange, true))}
	if (stages > 8) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeCertificateVerify, true))}

	return conn
}
func createOneway(stages int) Connection {
	conn := NewConnection("localhost", "localhost:443")

	if (stages > 0) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeClientHello, true))}
	if (stages > 1) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeServerHello, false))}
	if (stages > 2) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeCertificate, false))}
	if (stages > 3) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeServerKeyExchange, false))}
	if (stages > 4) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeServerHelloDone, false))}
	if (stages > 5) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeClientKeyExchange, true))}

	return conn
}

func CreateTestData() list.List {
	retval :=  list.List{}

	retval.PushBack(createMutual(20).DetectProblem(close_notify))
	retval.PushBack(createOneway(20).DetectProblem(close_notify))
	retval.PushBack(createOneway(5).DetectProblem(unknown_ca))
	retval.PushBack(createOneway(5).DetectProblem(certificate_expired))
	retval.PushBack(createOneway(5).DetectProblem(certificate_revoked))
	retval.PushBack(createOneway(5).DetectProblem(unknown_ca))
	retval.PushBack(createMutual(11).DetectProblem(unknown_ca))
	retval.PushBack(createOneway(5).DetectProblem(unknown_ca))

	return retval
}
