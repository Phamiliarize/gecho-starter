package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Phamiliarize/gecho-clean-starter/http/application"
	"github.com/Phamiliarize/gecho-clean-starter/service"
	"github.com/labstack/echo/v4"
)

func GetOneBook(app *application.Application) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		uid := uint32(id)
		if err != nil {
			fmt.Println("Do something about this error...")
		}

		validbook, err := service.GetOneBook(&uid)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		book, err := app.Repo.Book.One(&validbook.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong. Please try again.")
		}

		return c.JSON(http.StatusOK, &book)
	}
}
