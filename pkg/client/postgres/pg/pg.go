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

// ScanOne executes a query and scans the result into the provided destination.
func (p *pg) ScanOne(ctx context.Context, dest any, q postgres.Query, args ...any) error {
	row, err := p.Query(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, row)
}

// ScanAll executes a query and scans all the results into the provided destination.
func (p *pg) ScanAll(ctx context.Context, dest any, q postgres.Query, args ...any) error {
	rows, err := p.Query(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}

// Exec executes a query without returning any rows.
func (p *pg) Exec(ctx context.Context, q postgres.Query, args ...any) (pgconn.CommandTag, error) {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Exec(ctx, q.QueryRaw, args...)
	}

	return p.dbc.Exec(ctx, q.QueryRaw, args...)
}

// Query executes a query and returns the resulting rows.
func (p *pg) Query(ctx context.Context, q postgres.Query, args ...any) (pgx.Rows, error) {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Query(ctx, q.QueryRaw, args...)
	}

	return p.dbc.Query(ctx, q.QueryRaw, args...)
}

// QueryRow executes a query and returns a single row.
func (p *pg) QueryRow(ctx context.Context, q postgres.Query, args ...any) pgx.Row {
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

func logQuery(ctx context.Context, q postgres.Query, args ...any) {
	if slog.Default().Enabled(ctx, slog.LevelDebug) {
		prettyQuery := prettier.Pretty(q.QueryRaw, prettier.PlaceholderDollar, args...)
		
		slog.DebugContext(
			ctx,
			"sql",
			slog.String("query", prettyQuery),
		)
	}
}
