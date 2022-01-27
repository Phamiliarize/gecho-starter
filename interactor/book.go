package interactor

import (
	"github.com/Phamiliarize/gecho-clean-starter/entity"
	"github.com/Phamiliarize/gecho-clean-starter/repository"
	"github.com/Phamiliarize/gecho-clean-starter/util"
)

type BookInput struct {
	ID uint32 `param:"id"`
}

type BookOutput struct {
	entity.Book
}

func BookInteractor(request *BookInput) (*BookOutput, error) {
	var response BookOutput

	var book entity.Book
	book.ID = request.ID

	result, err := repository.Book(&book)

	if err == nil {
		response.ID = result.ID
		response.Name = result.Name
		response.Read = result.Read
	}

	return &response, err
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

func BookCollectionInteractor(request *BookCollectionInput) (*BookCollectionOutput, error) {
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

	items, count, cursor, err := repository.BookList(&requestCursor, request.Limit)
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
