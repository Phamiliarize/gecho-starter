package handler

import (
	"net/http"

	"github.com/Phamiliarize/gecho-clean-starter/repository"
	"github.com/Phamiliarize/gecho-clean-starter/service"
	"github.com/labstack/echo/v4"
)

type GetRequest struct {
	ID uint32 `param:"id"`
}

type GetResponse struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Read bool   `json:"read"`
}

func Get(c echo.Context) error {
	var request GetRequest
	c.Bind(&request)

	var repo repository.BookRepository
	repo = &repository.Book{}

	input := service.BookInput{ID: request.ID}

	book, err := service.Book(&input, repo)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "An internal server error has occurred. Please try again!")
	}

	response := GetResponse{ID: book.ID, Name: book.Name, Read: book.Read}

	return c.JSON(http.StatusOK, &response)
}

type GetListRequest struct {
	Limit     int    `query:"limit"`
	NextToken string `query:"nextToken"`
}

type GetListResponse struct {
	Count     int           `json:"count"`
	NextToken string        `json:"nextToken"`
	Items     []GetResponse `json:"items"`
}

func GetList(c echo.Context) error {
	var request GetListRequest
	c.Bind(&request)

	input := service.BookCollectionInput{Limit: request.Limit, NextToken: request.NextToken}

	var repo repository.BookRepository
	repo = &repository.Book{}

	bookCollection, err := service.BookCollection(&input, repo)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "An internal server error has occurred. Please try again!")
	}

	response := GetListResponse{Count: bookCollection.Count, NextToken: bookCollection.NextToken, Items: make([]GetResponse, len(bookCollection.Items))}

	for i, entity := range bookCollection.Items {
		response.Items[i] = GetResponse{
			ID:   entity.ID,
			Name: entity.Name,
			Read: entity.Read,
		}
	}

	return c.JSON(http.StatusOK, &response)
}
