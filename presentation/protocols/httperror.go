package protocols

import (
	"errors"
	"fmt"
)

type HttpError struct {
	StatusCode int
	Err        error
}

func (r *HttpError) Error() string {
	return fmt.Sprintf("[STATUS %d]: %v", r.StatusCode, r.Err)
}

func NewHttpError(message string, statusCode int) *HttpError {
	return &HttpError{
		Err:        errors.New(message),
		StatusCode: statusCode,
	}
}
