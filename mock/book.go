package mock

import (
	"errors"

	"github.com/Phamiliarize/gecho-clean-starter/entity"
)

type BookRepositoryMock struct{}

func (r BookRepositoryMock) FindByID(requestBook *entity.Book) (entity.Book, error) {
	switch requestBook.ID {
	case 1:
		return entity.Book{ID: 1, Name: "Test", Read: true}, nil
	default:
		return entity.Book{}, errors.New("Not found.") // For this level of testing we do not need exact errors, just an error
	}
}

func (r BookRepositoryMock) All(requestCursor *entity.Book, limit int) (entity.Books, int, entity.Book, error) {
	var cursor entity.Book
	var items entity.Books

	amount := 20
	if limit > amount {
		limit = amount
	}

	// Populate fake data
	books := make(entity.Books, amount)
	for i := 0; i < amount; i++ {
		books[i].ID = uint32(i + 1)
		books[i].Name = "Test"
	}

	count := len(books)

	// NextToken Present
	items = books[requestCursor.ID:]

	// Determine if we should be delivering a NextToken
	// Cut out non-relevant data
	items = items[:int(requestCursor.ID)]
	items = items[:limit]

	// Check if we need to return a cursor
	if items[len(items)-1].ID < books[amount-1].ID {
		cursor = items[len(items)-1]
	}

	return items, count, cursor, nil
}
