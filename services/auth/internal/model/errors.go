package model

import "errors"

var (
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

	// ErrTimeConverting is used for general time converting error.
	ErrTimeConverting = errors.New("time convert error")

	// ErrPasswordHashing is used for general password hashing error
	ErrPasswordHashing = errors.New("password hashing")

	// ErrInvalidCredentials is used when an invalid password is provided.
	ErrInvalidCredentials = errors.New("invalid credentials")

	// ErrInvalidToken is used when an invalid token is provided.
	ErrInvalidToken = errors.New("invalid token")
)
