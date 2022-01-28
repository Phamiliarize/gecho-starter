package book

import (
	"reflect"
	"testing"

	"github.com/Phamiliarize/gecho-clean-starter/entity"
	"github.com/Phamiliarize/gecho-clean-starter/repository"
)

func TestBook_BookInteractor(t *testing.T) {
	var repo repository.BookRepository
	repo = repository.BookRepositoryMock{}

	input := BookInput{ID: 1}

	result, err := BookInteractor(&input, repo)
	if err != nil {
		t.Errorf("Unexpected error, want: nil, got: '%v'.", err)
	}

	test := &BookOutput{entity.Book{ID: 1, Name: "Test", Read: true}}
	if !reflect.DeepEqual(test, result) {
		t.Errorf("Mismatched value, want: %v, got: %v.", result, test)
	}
}

func TestBook_BookInteractor_ForwardsError(t *testing.T) {
	var repo repository.BookRepository
	repo = repository.BookRepositoryMock{}

	input := BookInput{ID: 3}

	_, err := BookInteractor(&input, repo)
	if err == nil {
		t.Errorf("Expected error, want: 'Not found.', got: '%v'.", err)
	}
}
