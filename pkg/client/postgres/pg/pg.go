package pg

import (
	"context"
	"log/slog"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	postgres "github.com/alisher-baizhumanov/chat-microservices/pkg/client/postgres"
	"github.com/alisher-baizhumanov/chat-microservices/pkg/client/postgres/prettier"
)

type key string

const (
	// TxKey is a constant key used to store transaction objects in context.
	TxKey key = "tx"
)

type pg struct {
	dbc *pgxpool.Pool
}

// NewDB initializes a new pg instance with the given connection pool.
func NewDB(dbc *pgxpool.Pool) postgres.DB {
	return &pg{
		dbc: dbc,
	}
}

// ScanOneContext executes a query and scans the result into the provided destination.
func (p *pg) ScanOneContext(ctx context.Context, dest interface{}, q postgres.Query, args ...interface{}) error {
	logQuery(ctx, q, args...)

	row, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, row)
}

// ScanAllContext executes a query and scans all the results into the provided destination.
func (p *pg) ScanAllContext(ctx context.Context, dest interface{}, q postgres.Query, args ...interface{}) error {
	logQuery(ctx, q, args...)

	rows, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}

// ExecContext executes a query without returning any rows.
func (p *pg) ExecContext(ctx context.Context, q postgres.Query, args ...interface{}) (pgconn.CommandTag, error) {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Exec(ctx, q.QueryRaw, args...)
	}

	return p.dbc.Exec(ctx, q.QueryRaw, args...)
}

// QueryContext executes a query and returns the resulting rows.
func (p *pg) QueryContext(ctx context.Context, q postgres.Query, args ...interface{}) (pgx.Rows, error) {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Query(ctx, q.QueryRaw, args...)
	}

	return p.dbc.Query(ctx, q.QueryRaw, args...)
}

// QueryRowContext executes a query and returns a single row.
func (p *pg) QueryRowContext(ctx context.Context, q postgres.Query, args ...interface{}) pgx.Row {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.QueryRow(ctx, q.QueryRaw, args...)
	}

	return p.dbc.QueryRow(ctx, q.QueryRaw, args...)
}

// BeginTx begins a new transaction with the given options.
func (p *pg) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.dbc.BeginTx(ctx, txOptions)
}

// Ping checks the connection to the database.
func (p *pg) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}

// Close closes the database connection pool.
func (p *pg) Close() {
	p.dbc.Close()
}

// MakeContextTx returns a new context with the given transaction added to it.
func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}

func logQuery(ctx context.Context, q postgres.Query, args ...interface{}) {
	if slog.Default().Enabled(ctx, slog.LevelDebug) {
		prettyQuery := prettier.Pretty(q.QueryRaw, prettier.PlaceholderDollar, args...)
		slog.DebugContext(
			ctx,
			"sql",
			slog.String("query", prettyQuery),
		)
	}
}
