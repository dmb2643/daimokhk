package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func openDB(ctx context.Context, cn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, cn)
	if err != nil {
		return nil, fmt.Errorf("create PostgreSQL pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("Ping PostgreSQL: %w", err)
	}

	return pool, nil
}
