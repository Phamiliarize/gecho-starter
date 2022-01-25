package service

import (
	"github.com/Phamiliarize/gecho-clean-starter/entity"
)

type BookService interface {
	GetOneBook(id *uint32) (entity.Book, error)
}

func GetOneBook(id *uint32) (entity.Book, error) {
	//Do stuff with book
	book := entity.Book{ID: *id}
	err := book.Validate()

	return book, err
}
