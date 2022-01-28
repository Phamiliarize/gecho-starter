package book

import (
	"net/http"

	interactor "github.com/Phamiliarize/gecho-clean-starter/interactor/book"
	"github.com/Phamiliarize/gecho-clean-starter/repository"
	"github.com/labstack/echo/v4"
)

type GetBookListRequest struct {
	Limit     int    `query:"limit"`
	NextToken string `query:"nextToken"`
}

type BooksResponse []GetBookResponse

type GetBookListResponse struct {
	Count     int           `json:"count"`
	NextToken string        `json:"nextToken"`
	Items     BooksResponse `json:"items"`
}

func GetBookListController(c echo.Context) error {
	var request GetBookListRequest
	c.Bind(&request)

	input := interactor.BookCollectionInput{Limit: request.Limit, NextToken: request.NextToken}

	var repo repository.BookRepository
	repo = &repository.BookRepo{}

	bookCollection, err := interactor.BookCollectionInteractor(&input, repo)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "An internal server error has occurred. Please try again!")
	}

	response := GetBookListResponse{Count: bookCollection.Count, NextToken: bookCollection.NextToken, Items: make(BooksResponse, len(bookCollection.Items))}

	for i, entity := range bookCollection.Items {
		response.Items[i] = GetBookResponse{
			ID:   entity.ID,
			Name: entity.Name,
			Read: entity.Read,
		}
	}

	return c.JSON(http.StatusOK, &response)
}
