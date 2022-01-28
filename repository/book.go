package repository

import (
	"context"

	"github.com/Phamiliarize/gecho-clean-starter/entity"
	"github.com/georgysavva/scany/pgxscan"
)

// We use an interface so we can easily mock the Repository for testing
type BookRepository interface {
	FindByID(requestBook *entity.Book) (entity.Book, error)
	All(requestCursor *entity.Book, limit int) (entity.Books, int, entity.Book, error)
}

type BookRepo struct{}

// Postgres
// Returns a single book utilizing the ID
func (r BookRepo) FindByID(requestBook *entity.Book) (entity.Book, error) {
	var book entity.Book

	err := pgxscan.Get(context.Background(), pgPool, &book, `SELECT * FROM "book" WHERE id = $1`, requestBook.ID)

	return book, err
}

// Returns many books
func (r BookRepo) All(requestCursor *entity.Book, limit int) (entity.Books, int, entity.Book, error) {
	var books entity.Books
	var count int
	var cursor entity.Book

	// Presence of the +1 canary item tells us there is more data.
	canaryLimit := limit + 1

	// Start a transaction to ensure row count
	tx, err := pgPool.Begin(context.Background())
	if err != nil {
		return books, count, cursor, err
	}

	defer tx.Rollback(context.Background())

	err = pgxscan.Select(context.Background(), tx, &books, `SELECT * FROM "book" WHERE id > $1 LIMIT $2`, requestCursor.ID, canaryLimit)
	if err != nil {
		return books, count, cursor, err
	}

	err = tx.QueryRow(context.Background(), `SELECT COUNT(*) from "book"`).Scan(&count)
	if err != nil {
		return books, count, cursor, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return books, count, cursor, err
	}

	// Clean up the canary and set the cursor
	resultLength := len(books)
	if resultLength > limit {
		cursor.ID = books[resultLength-2].ID // Set the cursor
		books = books[:resultLength-1]       // Remove the canary item
	}

	return books, count, cursor, err
}
