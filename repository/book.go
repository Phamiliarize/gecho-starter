package repository

import (
	"context"

	"github.com/Phamiliarize/gecho-clean-starter/entity"
	"github.com/georgysavva/scany/pgxscan"
)

// Postgres
func FindOne(requestBook *entity.Book) (entity.Book, error) {
	var book entity.Book

	err := pgxscan.Get(context.Background(), pgPool, &book, `SELECT * FROM "book" WHERE id = $1`, requestBook.ID)

	return book, err
}
