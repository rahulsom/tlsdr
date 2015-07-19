package main

import (
	"errors"
	"github.com/rahulsom/TLSHandshakeDecoder"
	"log"
)

const (
	close_notify                int = 0
	unexpected_message          int = 10
	bad_record_mac              int = 20
	decryption_failed_RESERVED  int = 21
	record_overflow             int = 22
	decompression_failure       int = 30
	handshake_failure           int = 40
	no_certificate_RESERVED     int = 41
	bad_certificate             int = 42
	unsupported_certificate     int = 43
	certificate_revoked         int = 44
	certificate_expired         int = 45
	certificate_unknown         int = 46
	illegal_parameter           int = 47
	unknown_ca                  int = 48
	access_denied               int = 49
	decode_error                int = 50
	decrypt_error               int = 51
	export_restriction_RESERVED int = 60
	protocol_version            int = 70
	insufficient_security       int = 71
	internal_error              int = 80
	user_canceled               int = 90
	no_renegotiation            int = 100
	unsupported_extension       int = 110
	UNKNOWN                     int = 255
)

func (connection *Connection) DetectProblem(alert int) *Connection {
	DetectProblem(connection, alert)
	return connection
}

func DetectProblem(connection *Connection, alert int) {
	if alert == close_notify {
		process_close_notify(connection)
	} else {
		switch alert {
		case unexpected_message:
			process_unexpected_message(connection)
		case bad_record_mac:
			process_bad_record_mac(connection)
		case decryption_failed_RESERVED:
			process_decryption_failed_RESERVED(connection)
		case record_overflow:
			process_record_overflow(connection)
		case decompression_failure:
			process_decompression_failure(connection)
		case handshake_failure:
			process_handshake_failure(connection)
		case no_certificate_RESERVED:
			process_no_certificate_RESERVED(connection)
		case bad_certificate:
			process_bad_certificate(connection)
		case unsupported_certificate:
			process_unsupported_certificate(connection)
		case certificate_revoked:
			process_certificate_revoked(connection)
		case certificate_expired:
			process_certificate_expired(connection)
		case certificate_unknown:
			process_certificate_unknown(connection)
		case illegal_parameter:
			process_illegal_parameter(connection)
		case unknown_ca:
			process_unknown_ca(connection)
		case access_denied:
			process_access_denied(connection)
		case decode_error:
			process_decode_error(connection)
		case decrypt_error:
			process_decrypt_error(connection)
		case export_restriction_RESERVED:
			process_export_restriction_RESERVED(connection)
		case protocol_version:
			process_protocol_version(connection)
		case insufficient_security:
			process_insufficient_security(connection)
		case internal_error:
			process_internal_error(connection)
		case user_canceled:
			process_user_canceled(connection)
		case no_renegotiation:
			process_no_renegotiation(connection)
		case unsupported_extension:
			process_unsupported_extension(connection) /* new */
		case UNKNOWN:
			process_UNKNOWN(connection)
		}
	}
}

func findLastEvent(connection *Connection, code uint8) (*Event, error) {
	newCode := int(code)
	for e := connection.Events.Back(); e != nil; e = e.Prev() {
		// do something with e.Value
		if event, ok := e.Value.(*Event); ok {
			if event.Code == newCode {
				return event, nil
			}
		}
	}
	return nil, errors.New("Not found")
}
func process_close_notify(connection *Connection) {
	log.Println("close_notify is good!")
}
func process_unexpected_message(connection *Connection)         { log.Panicf("Not implemented") }
func process_bad_record_mac(connection *Connection)             { log.Panicf("Not implemented") }
func process_decryption_failed_RESERVED(connection *Connection) { log.Panicf("Not implemented") }
func process_record_overflow(connection *Connection)            { log.Panicf("Not implemented") }
func process_decompression_failure(connection *Connection)      { log.Panicf("Not implemented") }
func process_handshake_failure(connection *Connection)          { log.Panicf("Not implemented") }
func process_no_certificate_RESERVED(connection *Connection)    { log.Panicf("Not implemented") }
func process_bad_certificate(connection *Connection) {
	event, err := findLastEvent(connection, TLSHandshakeDecoder.HandshakeTypeCertificate)
	if err == nil {
		event.Success = false
	} else {
		log.Panicf("Didn't find event")
	}
	connection.Success = false
	connection.FailedReason = "The certificate was bad"
	connection.Recommendations.PushBack("Try matching the CN to the hostname")
}
func process_unsupported_certificate(connection *Connection) { log.Panicf("Not implemented") }
func process_certificate_revoked(connection *Connection) {
	event, err := findLastEvent(connection, TLSHandshakeDecoder.HandshakeTypeCertificate)
	if err == nil {
		event.Success = false
	} else {
		log.Panicf("Didn't find event")
	}
	connection.Success = false
	connection.FailedReason = "The certificate is revoked"
	connection.Recommendations.PushBack("Try getting a new certificate")
}
func process_certificate_expired(connection *Connection) {
	event, err := findLastEvent(connection, TLSHandshakeDecoder.HandshakeTypeCertificate)
	if err == nil {
		event.Success = false
	} else {
		log.Panicf("Didn't find event")
	}
	connection.Success = false
	connection.FailedReason = "The certificate is expired"
	connection.Recommendations.PushBack("Try getting a new certificate")
}
func process_certificate_unknown(connection *Connection) {
	log.Panicf("Not implemented")
}
func process_illegal_parameter(connection *Connection) { log.Panicf("Not implemented") }
func process_unknown_ca(connection *Connection) {
	event, err := findLastEvent(connection, TLSHandshakeDecoder.HandshakeTypeCertificate)
	if err == nil {
		event.Success = false
	} else {
		log.Panicf("Didn't find event")
	}
	connection.Success = false
	connection.FailedReason = "The CA is unknown"
	connection.Recommendations.PushBack("Try getting a certificate from a trusted CA")
	connection.Recommendations.PushBack("Try adding the CA of the issuer to the trust store")
}
func process_access_denied(connection *Connection) { log.Panicf("Not implemented") }
func process_decode_error(connection *Connection)  { log.Panicf("Not implemented") }
func process_decrypt_error(connection *Connection) {
	event, err := findLastEvent(connection, TLSHandshakeDecoder.HandshakeTypeCertificate)
	if err == nil {
		event.Success = false
	} else {
		log.Panicf("Didn't find event")
	}
	connection.Success = false
	connection.FailedReason = "Decrypting the traffic failed"
	connection.Recommendations.PushBack("Verify that the Certificate and private key match")
}
func process_export_restriction_RESERVED(connection *Connection) { log.Panicf("Not implemented") }
func process_protocol_version(connection *Connection)            { log.Panicf("Not implemented") }
func process_insufficient_security(connection *Connection)       { log.Panicf("Not implemented") }
func process_internal_error(connection *Connection)              { log.Panicf("Not implemented") }
func process_user_canceled(connection *Connection)               { log.Panicf("Not implemented") }
func process_no_renegotiation(connection *Connection)            { log.Panicf("Not implemented") }
func process_unsupported_extension(connection *Connection)       { log.Panicf("Not implemented") } /* new */
func process_UNKNOWN(connection *Connection)                     { log.Println("UNKNOWN not implemented") }
