package repository

import "github.com/jackc/pgx/v4/pgxpool"

type PostgresConn struct {
	DBC *pgxpool.Pool
}
