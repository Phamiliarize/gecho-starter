package handler

import (
	"net/http"

	"github.com/Phamiliarize/gecho-clean-starter/interactor"
	"github.com/labstack/echo/v4"
)

func GetBookListHandler(c echo.Context) error {
	var request interactor.GetBookCollectionRequest
	c.Bind(&request)

	response, err := interactor.GetBookCollectionInteractor(&request)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "An internal server error has occurred. Please try again!")
	}

	return c.JSON(http.StatusOK, response)
}

func GetBookHandler(c echo.Context) error {
	var request interactor.GetBookRequest
	c.Bind(&request)

	response, err := interactor.GetBookInteractor(&request)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "An internal server error has occurred. Please try again!")
	}

	return c.JSON(http.StatusOK, response)
}
