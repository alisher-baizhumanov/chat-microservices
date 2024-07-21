package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	db "github.com/alisher-baizhumanov/chat-microservices/pkg/client/postgres"
)

type pgClient struct {
	masterDBC db.DB
}

// New initializes a new PostgreSQL client with the given data source name (DSN).
func New(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	return &pgClient{
		masterDBC: &pg{dbc: dbc},
	}, nil
}

// DB returns the master database connection.
func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

// Close closes the master database connection.
func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
