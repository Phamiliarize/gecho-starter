package service

import (
	"github.com/Phamiliarize/gecho-clean-starter/entity"
	"github.com/Phamiliarize/gecho-clean-starter/repository"
	"github.com/Phamiliarize/gecho-clean-starter/util"
)

type BookService interface {
	BookCollection(request *BookCollectionInput) (*BookCollectionOutput, error)
	Book(request *BookInput) (*BookOutput, error)
}

type BookServiceImplement struct {
	Repo repository.BookRepository
}

type BookCollectionInput struct {
	Limit     int
	NextToken string
}

type BookCollectionOutput struct {
	Count     int
	NextToken string
	Items     entity.Books
}

func (d BookServiceImplement) BookCollection(request *BookCollectionInput) (*BookCollectionOutput, error) {
	var response BookCollectionOutput
	var requestCursor entity.Book

	// Default is 0
	if request.NextToken != "" {
		// Convert nextToken to requestCursor
		decoded, err := util.Uint32FromB64(request.NextToken)
		if err != nil {
			return &response, err
		}

		requestCursor.ID = decoded
	}

	if request.Limit <= 0 {
		request.Limit = 10
	}

	items, count, cursor, err := d.Repo.All(&requestCursor, request.Limit)
	if err != nil {
		return &response, err
	}

	// Handle Base64 Conversion
	response.Count = count
	if cursor.ID > 0 {
		response.NextToken = util.B64FromUint32(cursor.ID)
	}
	response.Items = items

	return &response, err
}

type BookInput struct {
	ID uint32 `param:"id"`
}

type BookOutput struct {
	entity.Book
}

func (d BookServiceImplement) Book(request *BookInput) (*BookOutput, error) {
	var response BookOutput

	var book entity.Book
	book.ID = request.ID

	result, err := d.Repo.FindByID(&book)

	if err == nil {
		response.ID = result.ID
		response.Name = result.Name
		response.Read = result.Read
	}

	return &response, err
}
