package reflect

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrNotAuthenticated = errors.New("not authenticated.")

	ErrConnectionFailed = errors.New("could not contact host.")

	ErrInternal = errors.New("Internal error on Reflect's end.")
)

type ReflectError struct {
	// The status code that was returned by the service.
	StatusCode int

	// The error message (if any) that was returned by the service.
	Message string

	// Internal message used for calling Errors()
	internal string
}

func (re *ReflectError) Error() string {
	return re.internal
}

func NewError(code int, message string, internal string) *ReflectError {
	re := new(ReflectError)
	re.StatusCode = code
	re.Message = message
	re.internal = internal
	return re
}

func NewErrorFromResponse(req *http.Response) *ReflectError {
	var message string

	// Decode the body, it should be a generatlized error.
	var reflectError struct {
		Error string `json:"error"`
	}

	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&reflectError)

	// If something bad happened here, we'll se that as the internal message.
	if err != nil {
		logError("Deserializing error message failed. %v", err)
		message = "(unknown)"
	} else {
		message = reflectError.Error
	}

	return NewError(req.StatusCode, message, "Unkown error from Reflect.")
}
