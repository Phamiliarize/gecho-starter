package handler

import (
	"net/http"

	"github.com/Phamiliarize/gecho-clean-starter/service"
	"github.com/labstack/echo/v4"
)

type BookHandler interface {
	BookGet(c echo.Context) error
	BookGetList(c echo.Context) error
}

type BookHandlerImplement struct {
	Service service.BookService
}

type BookGetRequest struct {
	ID uint32 `param:"id"`
}

type BookGetResponse struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Read bool   `json:"read"`
}

func (d BookHandlerImplement) BookGet(c echo.Context) error {
	var request BookGetRequest
	c.Bind(&request)

	input := service.BookInput{ID: request.ID}

	book, err := d.Service.Book(&input)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "An internal server error has occurred. Please try again!")
	}

	response := BookGetResponse{ID: book.ID, Name: book.Name, Read: book.Read}

	return c.JSON(http.StatusOK, &response)
}

type BookGetListRequest struct {
	Limit     int    `query:"limit"`
	NextToken string `query:"nextToken"`
}

type BookGetListResponse struct {
	Count     int               `json:"count"`
	NextToken string            `json:"nextToken"`
	Items     []BookGetResponse `json:"items"`
}

func (d BookHandlerImplement) BookGetList(c echo.Context) error {
	var request BookGetListRequest
	c.Bind(&request)

	input := service.BookCollectionInput{Limit: request.Limit, NextToken: request.NextToken}

	bookCollection, err := d.Service.BookCollection(&input)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "An internal server error has occurred. Please try again!")
	}

	response := BookGetListResponse{Count: bookCollection.Count, NextToken: bookCollection.NextToken, Items: make([]BookGetResponse, len(bookCollection.Items))}

	for i, entity := range bookCollection.Items {
		response.Items[i] = BookGetResponse{
			ID:   entity.ID,
			Name: entity.Name,
			Read: entity.Read,
		}
	}

	return c.JSON(http.StatusOK, &response)
}
