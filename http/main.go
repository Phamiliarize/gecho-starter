package main

import (
	"context"
	"log"
	"os"

	"github.com/Phamiliarize/gecho-clean-starter/app"
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
		log.Fatal("Error initializing database connection pool.")
	}

	defer dbPool.Close()

	// Instance App/Dependencies
	app := app.InitializeApp(dbPool)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/book", app.Handler.Book.BookGetList)
	e.GET("/book/:id", app.Handler.Book.BookGet)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
