package application

import (
	"github.com/Phamiliarize/gecho-clean-starter/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Applications is used to inject certain environment dependencies to the handler
type Application struct {
	Repo Repository
}

type Repository struct {
	Book repository.BookRepository
}

func InitializeApp(dbc *pgxpool.Pool) Application {
	return Application{
		Repo: Repository{
			Book: repository.PostgresConn{DBC: dbc},
		},
	}
}
