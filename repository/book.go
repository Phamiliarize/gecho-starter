package repository

import (
	"context"

	"github.com/Phamiliarize/gecho-clean-starter/entity"
	"github.com/georgysavva/scany/pgxscan"
)

type BookRepository interface {
	One(id *uint32) (entity.Book, error)
}

// Postgres
func (p PostgresConn) One(id *uint32) (entity.Book, error) {
	var book entity.Book

	err := pgxscan.Get(context.Background(), p.DBC, &book, `SELECT * FROM "book" WHERE id = $1`, id)

	return book, err
}

// In-memory/File etc could also go here and be a part of the interface. It's storage agnostic
