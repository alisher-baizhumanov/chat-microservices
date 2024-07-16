package model

import "errors"

var (
	// ErrCanNotBeNil is used when a function argument is expected to be non-nil but is provided as nil.
	ErrCanNotBeNil = errors.New("argument can not be null")

	// ErrInvalidID is used when an invalid ID is provided, for example, an ID that does not exist or is incorrectly formatted.
	ErrInvalidID = errors.New("invalid id")

	// ErrDatabase is used for general database errors.
	ErrDatabase = errors.New("database error")

	// ErrNotFound is used when a requested resource is not found.
	ErrNotFound = errors.New("not found")
)
