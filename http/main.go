package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Phamiliarize/gecho-clean-starter/http/handler"
	"github.com/Phamiliarize/gecho-clean-starter/repository"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize DBC Pool
	err = repository.InitDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error initializing database connection pool.")
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/book", handler.GetBookListHandler)
	e.GET("/book/:id", handler.GetBookHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
