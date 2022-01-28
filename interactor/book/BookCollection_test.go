package book

import (
	"reflect"
	"testing"

	"github.com/Phamiliarize/gecho-clean-starter/entity"
	"github.com/Phamiliarize/gecho-clean-starter/mock"
	"github.com/Phamiliarize/gecho-clean-starter/repository"
)

func TestBook_BookCollectionInteractor_DefaultInput(t *testing.T) {
	var repo repository.BookRepository
	repo = mock.BookRepositoryMock{}

	input := BookCollectionInput{}

	result, err := BookCollectionInteractor(&input, repo)
	if err != nil {
		t.Errorf("Unexpected error, want: nil, got: '%v'.", err)
	}

	testItems := make(entity.Books, 10)
	for i := 0; i < 10; i++ {
		testItems[i].ID = uint32(i + 1)
		testItems[i].Name = "Test"
	}

	test := &BookCollectionOutput{Count: 20, NextToken: "MTA=", Items: testItems}
	if !reflect.DeepEqual(test, result) {
		t.Errorf("Mismatched value, want: %v, got: %v.", test, result)
	}
}

func TestBook_BookCollectionInteractor_CustomInput(t *testing.T) {
	var repo repository.BookRepository
	repo = mock.BookRepositoryMock{}

	input := BookCollectionInput{Limit: 10, NextToken: "MTA="}

	result, err := BookCollectionInteractor(&input, repo)
	if err != nil {
		t.Errorf("Unexpected error, want: nil, got: '%v'.", err)
	}

	testItems := make(entity.Books, 10)
	for i := 0; i < 10; i++ {
		testItems[i].ID = uint32(i + 11)
		testItems[i].Name = "Test"
	}

	test := &BookCollectionOutput{Count: 20, NextToken: "", Items: testItems}
	if !reflect.DeepEqual(test, result) {
		t.Errorf("Mismatched value, want: %v, got: %v.", test, result)
	}
}

func TestBook_BookCollectionInteractor_NoNext(t *testing.T) {
	var repo repository.BookRepository
	repo = mock.BookRepositoryMock{}

	input := BookCollectionInput{Limit: 20}

	result, err := BookCollectionInteractor(&input, repo)
	if err != nil {
		t.Errorf("Unexpected error, want: nil, got: '%v'.", err)
	}

	testItems := make(entity.Books, 20)
	for i := 0; i < 20; i++ {
		testItems[i].ID = uint32(i + 1)
		testItems[i].Name = "Test"
	}

	test := &BookCollectionOutput{Count: 20, NextToken: "", Items: testItems}
	if !reflect.DeepEqual(test, result) {
		t.Errorf("Mismatched value, want: %v, got: %v.", test, result)
	}
}

func TestBook_BookCollectionInteractor_BadNextTokenError(t *testing.T) {
	var repo repository.BookRepository
	repo = mock.BookRepositoryMock{}

	input := BookCollectionInput{NextToken: "ABC"}

	_, err := BookCollectionInteractor(&input, repo)
	if err == nil {
		t.Errorf("Expected error due to bad NextToken, got: '%v'.", err)
	}
}
