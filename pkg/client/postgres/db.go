package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Handler - a function that is executed within a transaction.
type Handler func(ctx context.Context) error

// Client - a client for interacting with the database.
type Client interface {
	DB() DB
	Close() error
}

// TxManager - a transaction manager that executes a user-specified handler within a transaction.
type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

// Query - a wrapper around a query that holds the name of the query and the raw query string.
// The query name is used for logging and can potentially be used elsewhere, such as for tracing.
type Query struct {
	Name     string
	QueryRaw string
}

// Transactor - an interface for working with transactions.
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

// SQLExecer - combines NamedExecer and QueryExecer.
type SQLExecer interface {
	NamedExecer
	QueryExecer
}

// NamedExecer - an interface for working with named queries using struct tags.
type NamedExecer interface {
	ScanOne(ctx context.Context, dest any, q Query, args ...any) error
	ScanAll(ctx context.Context, dest any, q Query, args ...any) error
}

// QueryExecer - an interface for working with regular queries.
type QueryExecer interface {
	Exec(ctx context.Context, q Query, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, q Query, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, q Query, args ...any) pgx.Row
}

// Pinger - an interface for checking the database connection.
type Pinger interface {
	Ping(ctx context.Context) error
}

// DB - an interface for interacting with the database.
type DB interface {
	SQLExecer
	Transactor
	Pinger
	Close()
}
