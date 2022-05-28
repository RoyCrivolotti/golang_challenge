package enums

import (
	"fmt"
	"strings"
)

type ErrorType int

const (
	NotFound ErrorType = iota + 1
	UnexpectedDatabaseError
	Forbidden
	BadRequest
	InternalServer
	InvalidData
)

var ErrorTypeSlice = []string{"NotFound", "UnexpectedDatabaseError", "InternalServerError", "Forbidden", "BadRequest", "InvalidData"}

func (s ErrorType) String() string {
	return ErrorTypeSlice[s-1]
}

func ErrorTypeFromString(value string) (error, ErrorType) {
	value = strings.ToUpper(value)
	for i := 0; i < len(ErrorTypeSlice); i++ {
		if ErrorTypeSlice[i] == value {
			return nil, ErrorType(i + 1)
		}
	}
	return fmt.Errorf("invalid error type: %s", value), 0
}
