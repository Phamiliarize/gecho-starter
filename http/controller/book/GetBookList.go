package book

import (
	"net/http"

	"github.com/Phamiliarize/gecho-clean-starter/entity"
	"github.com/Phamiliarize/gecho-clean-starter/interactor"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type GetBookListRequest struct {
	Limit     int    `query:"limit"`
	NextToken string `query:"nextToken"`
}

type GetBookListResponse struct {
	Count     int          `json:"count"`
	NextToken string       `json:"nextToken"`
	Items     entity.Books `json:"items"`
}

func GetBookListController(c echo.Context) error {
	var request GetBookListRequest
	c.Bind(&request)

	input := interactor.BookCollectionInput{Limit: request.Limit, NextToken: request.NextToken}

	bookCollection, err := interactor.BookCollectionInteractor(&input)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "An internal server error has occurred. Please try again!")
	}

	response := GetBookListResponse{Count: bookCollection.Count, NextToken: bookCollection.NextToken, Items: make(entity.Books, 0)}
	length := len(bookCollection.Items)
	if length > 0 {
		copier.Copy(&response.Items, bookCollection.Items)
	}

	return c.JSON(http.StatusOK, &response)
}
