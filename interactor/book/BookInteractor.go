package book

import (
	"github.com/Phamiliarize/gecho-clean-starter/entity"
	"github.com/Phamiliarize/gecho-clean-starter/repository"
)

type BookInput struct {
	ID uint32 `param:"id"`
}

type BookOutput struct {
	entity.Book
}

func BookInteractor(request *BookInput, repo repository.BookRepository) (*BookOutput, error) {
	var response BookOutput

	var book entity.Book
	book.ID = request.ID

	result, err := repo.FindByID(&book)

	if err == nil {
		response.ID = result.ID
		response.Name = result.Name
		response.Read = result.Read
	}

	return &response, err
}
