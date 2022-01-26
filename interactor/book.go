package interactor

import (
	"github.com/Phamiliarize/gecho-clean-starter/entity"
	"github.com/Phamiliarize/gecho-clean-starter/repository"
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

	result, err := repository.FindOne(&book)

	if err == nil {
		response.ID = result.ID
		response.Name = result.Name
		response.Read = result.Read
	}

	return &response, err
}
