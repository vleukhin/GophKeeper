package errs

import (
	"errors"
)

var (
	ErrWrongEmail           = errors.New("incorrect email given")
	ErrEmailAlreadyExists   = errors.New("given email already exists")
	ErrWrongCredentials     = errors.New("wrong credentials have been given")
	ErrTokenValidation      = errors.New("token validation error")
	ErrUnexpectedError      = errors.New("some unexpected error")
	ErrWrongOwnerOrNotFound = errors.New("wrong owner or not found")
)
