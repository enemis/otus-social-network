package app_error

import "net/http"

type HttpError struct {
	message       string
	field         string
	status        int
	originalError error
}

const SERVER_ERROR_MESSAGE = "Something went wrong"
const DEFAULT_ERROR_FIELD = "error"

func (e HttpError) Error() string {
	return e.message
}

func (e HttpError) Field() string {
	return e.field
}

func (e HttpError) Status() int {
	return e.status
}

func (e HttpError) OriginalError() string {
	if e.originalError != nil {
		return e.originalError.Error()
	}

	return e.message
}

func NewHttpError(originalError error, message, field string, status int) *HttpError {
	if message == "" {
		message = SERVER_ERROR_MESSAGE
	}

	if status == 0 {
		status = http.StatusInternalServerError
	}

	if field == "" {
		field = DEFAULT_ERROR_FIELD
	}

	return &HttpError{message: message, field: field, status: status, originalError: originalError}
}

func NewInternalServerError(originalError error) *HttpError {
	return NewHttpError(originalError, "", "", 0)
}
