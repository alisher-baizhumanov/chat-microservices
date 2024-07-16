package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewConnectionPool initializes a new PostgreSQL connection pool.
func NewConnectionPool(ctx context.Context, connectionString string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}

// CloseConnectionPool gracefully closes the given connection pool if it is not nil.
func CloseConnectionPool(pool *pgxpool.Pool) {
	if pool == nil {
		return
	}

	pool.Close()
}
