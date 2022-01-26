package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Postgres Connection Pool
var pgPool *pgxpool.Pool

// Initializes the connection pool and attempts a query
func InitDB(databaseUrl string) error {
	var err error

	pgPool, err = pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		return err
	}

	return pgPool.Ping(context.Background())
}
