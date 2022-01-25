package interactor

import (
	"fmt"
	"net/http"

	"github.com/Phamiliarize/gecho-clean-starter/entity"
	"github.com/Phamiliarize/gecho-clean-starter/http/application"
)

// type BookInteractor interface {
// 	GetBookInteractor(request *GetBookRequest) *GetBookResponse
// }

type Book struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Read bool   `json:"read"`
}

type GetBookRequest struct {
	ID uint32 `param:"id"`
}

type GetBookResponse struct {
	Status int
	Body   Book
}

func GetBookInteractor(request *GetBookRequest, app *application.Application) *GetBookResponse {
	var response GetBookResponse

	var book entity.Book
	book.ID = request.ID

	result, err := app.Repo.Book.FindOne(&book)
	if err != nil {
		fmt.Println("how should we handle errors")
	}

	response.Status = http.StatusOK
	response.Body.ID = result.ID
	response.Body.Name = result.Name
	response.Body.Read = result.Read

	return &response
}
