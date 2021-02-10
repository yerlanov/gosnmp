package exception

import "fmt"

const (
	OidConvertError  = "OID_CONVERT_ERROR"
	ParseError       = "PARSE_ERROR"
	SnmpConnectError = "SNMP_CONNECT_ERROR"
	SnmpGetError     = "SNMP_GET_ERROR"
	DecodeError      = "DECODE_ERROR"
	NOTFOUND         = "NOT_FOUND"
)

type Error struct {
	ErrorType    string
	ErrorMessage error
}

func (handleError *Error) Error() string {
	return fmt.Sprintf("%v: %s", handleError.ErrorType, handleError.ErrorMessage)
}
