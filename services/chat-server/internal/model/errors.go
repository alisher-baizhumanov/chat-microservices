package model

import "errors"

var (
	// ErrInvalidID is used when an invalid ID is provided, for example, an ID that does not exist or is incorrectly formatted.
	ErrInvalidID = errors.New("invalid id")

	// ErrDatabase is used for general database errors.
	ErrDatabase = errors.New("database error")

	// ErrNotFound is used when a requested resource is not found.
	ErrNotFound = errors.New("not found")

	// ErrGeneratingID is used when there is an error generating an ID.
	ErrGeneratingID = errors.New("error generating id")
)
