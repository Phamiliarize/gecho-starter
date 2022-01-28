package entity

// Represents a single book
type Book struct {
	ID   uint32 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Read bool   `json:"read" db:"read"`
}

// Represents a collection of books
type Books []Book
