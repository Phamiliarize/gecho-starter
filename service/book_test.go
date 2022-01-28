package service

import (
	"reflect"
	"testing"

	"github.com/Phamiliarize/gecho-clean-starter/entity"
	"github.com/Phamiliarize/gecho-clean-starter/repository"
)

func TestBook_Book(t *testing.T) {
	var service BookService
	service = BookServiceImplement{Repo: repository.BookRepositoryMock{}}

	input := BookInput{ID: 1}

	result, err := service.Book(&input)
	if err != nil {
		t.Errorf("Unexpected error, want: nil, got: '%v'.", err)
	}

	test := &BookOutput{entity.Book{ID: 1, Name: "Test", Read: true}}
	if !reflect.DeepEqual(test, result) {
		t.Errorf("Mismatched value, want: %v, got: %v.", result, test)
	}
}

func TestBook_Book_ForwardsError(t *testing.T) {
	var service BookService
	service = BookServiceImplement{Repo: repository.BookRepositoryMock{}}

	input := BookInput{ID: 3}

	_, err := service.Book(&input)
	if err == nil {
		t.Errorf("Expected error, want: 'Not found.', got: '%v'.", err)
	}
}

func TestBook_BookCollection_DefaultInput(t *testing.T) {
	var service BookService
	service = BookServiceImplement{Repo: repository.BookRepositoryMock{}}

	input := BookCollectionInput{}

	result, err := service.BookCollection(&input)
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

func TestBook_BookCollection_CustomInput(t *testing.T) {
	var service BookService
	service = BookServiceImplement{Repo: repository.BookRepositoryMock{}}

	input := BookCollectionInput{Limit: 10, NextToken: "MTA="}

	result, err := service.BookCollection(&input)
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

func TestBook_BookCollection_NoNext(t *testing.T) {
	var service BookService
	service = BookServiceImplement{Repo: repository.BookRepositoryMock{}}

	input := BookCollectionInput{Limit: 20}

	result, err := service.BookCollection(&input)
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

func TestBook_BookCollection_BadNextTokenError(t *testing.T) {
	var service BookService
	service = BookServiceImplement{Repo: repository.BookRepositoryMock{}}

	input := BookCollectionInput{NextToken: "ABC"}

	_, err := service.BookCollection(&input)
	if err == nil {
		t.Errorf("Expected error due to bad NextToken, got: '%v'.", err)
	}
}
