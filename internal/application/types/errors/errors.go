package types

import "errors"

// Error is a custom error type that implements the error interface

var (
	ErrAlreadyExists = errors.New("resource already exists")
)