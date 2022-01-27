package book

import (
	"net/http"

	"github.com/Phamiliarize/gecho-clean-starter/interactor"
	"github.com/labstack/echo/v4"
)

type GetBookRequest struct {
	ID uint32 `param:"id"`
}

type GetBookResponse struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Read bool   `json:"read"`
}

func GetBookController(c echo.Context) error {
	var request GetBookRequest
	c.Bind(&request)

	input := interactor.BookInput{ID: request.ID}

	book, err := interactor.BookInteractor(&input)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "An internal server error has occurred. Please try again!")
	}

	response := GetBookResponse{ID: book.ID, Name: book.Name, Read: book.Read}

	return c.JSON(http.StatusOK, &response)
}
