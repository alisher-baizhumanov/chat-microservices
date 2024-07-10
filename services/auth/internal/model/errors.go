package model

import "fmt"

var (
	// ErrCanNotBeNil is used when a function argument is expected to be non-nil but is provided as nil.
	ErrCanNotBeNil = fmt.Errorf("argument can not be null")

	// ErrInvalidID is used when an invalid ID is provided, for example, an ID that does not exist or is incorrectly formatted.
	ErrInvalidID = fmt.Errorf("invalid id")
)
