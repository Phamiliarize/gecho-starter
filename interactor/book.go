package interactor

import (
	b64 "encoding/base64"
	"strconv"

	"github.com/Phamiliarize/gecho-clean-starter/entity"
	"github.com/Phamiliarize/gecho-clean-starter/repository"
	"github.com/jinzhu/copier"
)

// type BookInteractor interface {
// 	GetBookInteractor(request *GetBookRequest) *GetBookResponse
// }

type GetBookRequest struct {
	ID uint32 `param:"id"`
}

type GetBookResponse struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Read bool   `json:"read"`
}

func GetBookInteractor(request *GetBookRequest) (*GetBookResponse, error) {
	var response GetBookResponse

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

type GetBookCollectionRequest struct {
	Limit     int    `query:"Limit"`
	NextToken string `query:"nextToken"`
}

type GetBookCollectionResponse struct {
	Count     int               `json:"count"`
	NextToken string            `json:"nextToken"`
	Items     []GetBookResponse `json:"items"`
}

func GetBookCollectionInteractor(request *GetBookCollectionRequest) (*GetBookCollectionResponse, error) {
	var response GetBookCollectionResponse
	var requestCursor entity.Book

	// Default is 0
	if request.NextToken != "" {
		// Convert nextToken to requestCursor
		byteDecoded, err := b64.StdEncoding.DecodeString(request.NextToken)
		if err != nil {
			return &response, err
		}

		decoded, err := strconv.ParseUint(string(byteDecoded), 10, 32)
		if err != nil {
			return &response, err
		}

		if decoded > 1 {
			requestCursor.ID = uint32(decoded)
		}
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
		response.NextToken = b64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(int64(cursor.ID), 10)))
	}
	copier.Copy(&response.Items, &items)

	return &response, err
}
