package errorx

import (
	"fmt"
	"golangchallenge/internal/constants/enums"
)

type Error interface {
	ClientMessage() string
	Error() string
	ErrorType() enums.ErrorType
	JSON() interface{}
}

type err struct {
	errorType     enums.ErrorType
	clientMessage string
	error         string
	thrownAt      string
}

//ClientMessage returns the clientMessage for the current error
func (e *err) ClientMessage() string {
	return e.clientMessage
}

//Error returns the error
func (e *err) Error() string {
	return fmt.Sprintf("Thrown at: %s - ClientMessage: %s - Error: %s", e.thrownAt, e.clientMessage, e.error)
}

//ErrorType returns the type of error
func (e *err) ErrorType() enums.ErrorType {
	return e.errorType
}

//JSON returns the Json codification for the current error
func (e *err) JSON() interface{} {
	return &struct {
		ErrorType enums.ErrorType `json:"error_type"`
		Message   string          `json:"clientMessage"`
	}{
		e.errorType,
		e.clientMessage,
	}
}

//NewNotFound creates an error configured to be a NotFound response
func NewNotFound(clientMsg string, error string, thrownAt string) Error {
	return &err{
		errorType:     enums.NotFound,
		clientMessage: clientMsg,
		error:         error,
		thrownAt:      thrownAt,
	}
}

//NewUnexpectedDatabaseError corresponds to an unexpected error thrown by the database wrapper
func NewUnexpectedDatabaseError(clientMsg string, error string, thrownAt string) Error {
	return &err{
		errorType:     enums.UnexpectedDatabaseError,
		clientMessage: clientMsg,
		error:         error,
		thrownAt:      thrownAt,
	}
}

//NewForbiddenError corresponds to a forbidden error
func NewForbiddenError(clientMsg string, error string, thrownAt string) Error {
	return &err{
		errorType:     enums.Forbidden,
		clientMessage: clientMsg,
		error:         error,
		thrownAt:      thrownAt,
	}
}

//NewBadRequestError corresponds to an issue with the request input
func NewBadRequestError(clientMsg string, error string, thrownAt string) Error {
	return &err{
		errorType:     enums.BadRequest,
		clientMessage: clientMsg,
		error:         error,
		thrownAt:      thrownAt,
	}
}

//NewInternalServerError corresponds to an unexpected error thrown by the database wrapper
func NewInternalServerError(clientMsg string, error string, thrownAt string) Error {
	return &err{
		errorType:     enums.InternalServer,
		clientMessage: clientMsg,
		error:         error,
		thrownAt:      thrownAt,
	}
}

//NewInvalidDataError corresponds to invalid data (probably from the incoming HTTP request's parameters)
func NewInvalidDataError(clientMsg string, error string, thrownAt string) Error {
	return &err{
		errorType:     enums.InvalidData,
		clientMessage: clientMsg,
		error:         error,
		thrownAt:      thrownAt,
	}
}
