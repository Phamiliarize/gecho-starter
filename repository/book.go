package repository

import (
	"context"

	"github.com/Phamiliarize/gecho-clean-starter/entity"
	"github.com/georgysavva/scany/pgxscan"
)

// By implementing through the interface and entity, we allow these methods to change as needed with no impact on interactors
type BookRepository interface {
	FindOne(book *entity.Book) (entity.Book, error)
}

// Postgres
func (p PostgresConn) FindOne(requestBook *entity.Book) (entity.Book, error) {
	var book entity.Book

	err := pgxscan.Get(context.Background(), p.DBC, &book, `SELECT * FROM "book" WHERE id = $1`, book.ID)

	return book, err
}

// Mongodb, an in-memory database, or even filesystem methods could go here, all obscured under one interface
