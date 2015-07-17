package main
import (
	"container/list"
	"github.com/certifyTian/TLSHandshakeDecoder"
)

func createMutual(stages int) Connection {
	conn := NewConnection("localhost", "localhost:443")

	if (stages > 0) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeClientHello))}
	if (stages > 1) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeServerHello))}
	if (stages > 2) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeCertificate))}
	if (stages > 3) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeServerKeyExchange))}
	if (stages > 4) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeCertificateRequest))}
	if (stages > 5) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeServerHelloDone))}
	if (stages > 6) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeCertificate))}
	if (stages > 7) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeClientKeyExchange))}
	if (stages > 8) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.HandshakeTypeCertificateVerify))}
	if (stages > 9) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.TypeChangeCypherSpec))}
	if (stages > 10) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.TypeHandshake))}
	if (stages > 11) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.TypeChangeCypherSpec))}
	if (stages > 12) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.TypeHandshake))}
	if (stages > 13) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.TypeApplicationData))}
	if (stages > 14) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.TypeApplicationData))}
	if (stages > 15) {conn.AddEvent(NewEvent(TLSHandshakeDecoder.TypeAlert))}

	return conn
}

func CreateTestData() list.List {
	retval :=  list.List{}

	conn := createMutual(15)
	DetectProblem(conn, close_notify)
	retval.PushBack(conn)

	return retval
}
