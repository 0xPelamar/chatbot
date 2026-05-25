package db

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPgPool(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		slog.Error("failed to parse the pg config", "error", err)
		return nil, err
	}
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		slog.Error("failed to connect to pgxpool", "error", err)
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		slog.Error("failed to ping pgxpool", "error", err)
		return nil, err
	}
	return pool, nil
}
