package model

import "errors"

var (
	// ErrCanNotBeNil is used when a function argument is expected to be non-nil but is provided as nil.
	ErrCanNotBeNil = errors.New("argument can not be null")

	// ErrInvalidID is used when an invalid ID is provided, for example, an ID that does not exist or is incorrectly formatted.
	ErrInvalidID = errors.New("invalid id")

	// ErrInvalidSQLQuery is returned when an SQL query is invalid.
	ErrInvalidSQLQuery = errors.New("invalid sql query")

	// ErrNonUniqueEmail is used when a provided email must be unique but is already in use.
	ErrNonUniqueEmail = errors.New("email must be unique")

	// ErrNonUniqueUsername is used when a provided username must be unique but is already in use.
	ErrNonUniqueUsername = errors.New("username must be unique")

	// ErrDatabase is used for general database errors.
	ErrDatabase = errors.New("database error")

	// ErrNotFound is used when a requested resource is not found.
	ErrNotFound = errors.New("not found")

	// ErrCache is used for general cache errors.
	ErrCache = errors.New("cache error")
)
