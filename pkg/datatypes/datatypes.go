package datatypes

// Book is a struct that represents a book
type Book struct {
	ID                int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title             string `gorm:"not null;unique" json:"title" validate:"required"`
	Author            string `json:"author" validate:"required"`
	Publisher         string `json:"publisher" validate:"required"`
	YearOfPublication int32  `json:"year_of_publication" validate:"required"`
}

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"not null;unique" json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`
}
