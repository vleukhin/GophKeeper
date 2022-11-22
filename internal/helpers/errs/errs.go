package errs

import (
	"errors"
)

var (
	ErrWrongCredentials     = errors.New("wrong credentials have been given")
	ErrTokenValidation      = errors.New("token validation error")
	ErrUnexpectedError      = errors.New("some unexpected error")
	ErrWrongOwnerOrNotFound = errors.New("wrong owner or not found")
)
