package handler

import (
	"net/http"

	"github.com/Phamiliarize/gecho-clean-starter/http/application"
	"github.com/Phamiliarize/gecho-clean-starter/interactor"
	"github.com/labstack/echo/v4"
)

func GetBookHandler(app *application.Application) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request interactor.GetBookRequest
		c.Bind(&request)

		response, err := interactor.GetBookInteractor(&request, app)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		// TODO: Is there really a difference between echo.NewHTTPError & c.JSON with an error code + body?
		//return echo.NewHTTPError(response.Staus, err)
		return c.JSON(http.StatusOK, response)
	}
}
