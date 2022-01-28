package app

import (
	"github.com/Phamiliarize/gecho-clean-starter/http/handler"
	"github.com/Phamiliarize/gecho-clean-starter/repository"
	"github.com/Phamiliarize/gecho-clean-starter/service"
	"github.com/jackc/pgx/v4/pgxpool"
)

// App is a dependency container for the api
type App struct {
	Handler Handler
}

type Handler struct {
	Book handler.BookHandler
}

func InitializeApp(dbPool *pgxpool.Pool) App {
	return App{
		Handler: Handler{
			Book: handler.BookHandlerImplement{
				Service: service.BookServiceImplement{Repo: repository.BookRepositoryImplement{PgPool: dbPool}},
			},
		},
	}
}
