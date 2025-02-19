package datatypes

// Book is a struct that represents a book
type Book struct {
	ID                int    `json:"id"`
	Title             string `json:"title" validate:"required"`
	Author            string `json:"author" validate:"required"`
	Publisher         string `json:"publisher" validate:"required"`
	YearOfPublication int    `json:"year_of_publication" validate:"required"`
}
