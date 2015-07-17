package main
import "log"

const (
	close_notify int8 = 0
	unexpected_message = 10
	bad_record_mac = 20
	decryption_failed_RESERVED = 21
	record_overflow = 22
	decompression_failure = 30
	handshake_failure = 40
	no_certificate_RESERVED = 41
	bad_certificate = 42
	unsupported_certificate = 43
	certificate_revoked = 44
	certificate_expired = 45
	certificate_unknown = 46
	illegal_parameter = 47
	unknown_ca = 48
	access_denied = 49
	decode_error = 50
	decrypt_error = 51
	export_restriction_RESERVED = 60
	protocol_version = 70
	insufficient_security = 71
	internal_error = 80
	user_canceled = 90
	no_renegotiation = 100
	unsupported_extension = 110
	UNKNOWN = 255
)

func DetectProblem(connection Connection, alert int) {
	switch alert {
	case close_notify: return process_close_notify(connection)
	case unexpected_message: return process_unexpected_message(connection)
	case bad_record_mac: return process_bad_record_mac(connection)
	case decryption_failed_RESERVED: return process_decryption_failed_RESERVED(connection)
	case record_overflow: return process_record_overflow(connection)
	case decompression_failure: return process_decompression_failure(connection)
	case handshake_failure: return process_handshake_failure(connection)
	case no_certificate_RESERVED: return process_no_certificate_RESERVED(connection)
	case bad_certificate: return process_bad_certificate(connection)
	case unsupported_certificate: return process_unsupported_certificate(connection)
	case certificate_revoked: return process_certificate_revoked(connection)
	case certificate_expired: return process_certificate_expired(connection)
	case certificate_unknown: return process_certificate_unknown(connection)
	case illegal_parameter: return process_illegal_parameter(connection)
	case unknown_ca: return process_unknown_ca(connection)
	case access_denied: return process_access_denied(connection)
	case decode_error: return process_decode_error(connection)
	case decrypt_error: return process_decrypt_error(connection)
	case export_restriction_RESERVED: return process_export_restriction_RESERVED(connection)
	case protocol_version: return process_protocol_version(connection)
	case insufficient_security: return process_insufficient_security(connection)
	case internal_error: return process_internal_error(connection)
	case user_canceled: return process_user_canceled(connection)
	case no_renegotiation: return process_no_renegotiation(connection)
	case unsupported_extension: return process_unsupported_extension(connection) /* new */
	case UNKNOWN: return process_UNKNOWN(connection)
	}
}

func process_close_notify(connection Connection) {
	log.Println("close_notify is good!")
}
func process_unexpected_message(connection Connection) { log.Panicf("Not implemented") }
func process_bad_record_mac(connection Connection) { log.Panicf("Not implemented") }
func process_decryption_failed_RESERVED(connection Connection) { log.Panicf("Not implemented") }
func process_record_overflow(connection Connection) { log.Panicf("Not implemented") }
func process_decompression_failure(connection Connection) { log.Panicf("Not implemented") }
func process_handshake_failure(connection Connection) { log.Panicf("Not implemented") }
func process_no_certificate_RESERVED(connection Connection) { log.Panicf("Not implemented") }
func process_bad_certificate(connection Connection) { log.Panicf("Not implemented") }
func process_unsupported_certificate(connection Connection) { log.Panicf("Not implemented") }
func process_certificate_revoked(connection Connection) {
	log.Panicf("Not implemented")
}
func process_certificate_expired(connection Connection) {
	log.Panicf("Not implemented")
}
func process_certificate_unknown(connection Connection) {
	log.Panicf("Not implemented")
}
func process_illegal_parameter(connection Connection) { log.Panicf("Not implemented") }
func process_unknown_ca(connection Connection) {
	log.Panicf("Not implemented")
}
func process_access_denied(connection Connection) { log.Panicf("Not implemented") }
func process_decode_error(connection Connection) { log.Panicf("Not implemented") }
func process_decrypt_error(connection Connection) {
	log.Panicf("Not implemented")
}
func process_export_restriction_RESERVED(connection Connection) { log.Panicf("Not implemented") }
func process_protocol_version(connection Connection) { log.Panicf("Not implemented") }
func process_insufficient_security(connection Connection) { log.Panicf("Not implemented") }
func process_internal_error(connection Connection) { log.Panicf("Not implemented") }
func process_user_canceled(connection Connection) { log.Panicf("Not implemented") }
func process_no_renegotiation(connection Connection) { log.Panicf("Not implemented") }
func process_unsupported_extension(connection Connection) { log.Panicf("Not implemented") } /* new */
func process_UNKNOWN(connection Connection) { log.Panicf("Not implemented") }

