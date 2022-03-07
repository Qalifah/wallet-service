package postgres

import (
	"context"
	"fmt"
	"github.com/Qalifah/grey-challenge/wallet/config"
	"github.com/jackc/pgx/v4"
)

func New(ctx context.Context, config *config.PostgresConfig) (*pgx.Conn, error) {
	const format = "postgres://%s:%s@%s:%s/%s?sslmode=disable"
	uri := fmt.Sprintf(format, config.Username, config.Password, config.Host, config.Port, config.Database)
	conn, err := pgx.Connect(ctx, uri)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func CreateTestDB(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://qali:qaliforshort@localhost:5432/grey")
}
