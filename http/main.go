package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Phamiliarize/gecho-clean-starter/http/application"
	"github.com/Phamiliarize/gecho-clean-starter/http/handler"
	"github.com/jackc/pgx/v4/pgxpool"
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
	dbPool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error initializing connection to DB")
	}

	defer dbPool.Close()

	// Initialize Application Environment

	app := application.InitializeApp(dbPool)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/book/:id", handler.GetBookHandler(&app))

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
