package entity

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Represents a single book
type Book struct {
	ID   uint32 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Read bool   `json:"read" db:"read"`
}

// Represents a collection of books
type Books []Book

// Validates a book (only looks at present values)
func (b Book) Validate() error {
	err := validation.ValidateStruct(&b,
		// Only name requires validation- rest are handled by type
		validation.Field(&b.Name, validation.Length(3, 128)),
	)

	return err
}
