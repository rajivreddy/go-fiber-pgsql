package handlers

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/rajivreddy/go-fiber-pgsql/pkg/datatypes"
	"github.com/rajivreddy/go-fiber-pgsql/pkg/postgres"
)

// Add Fiber Handlers here
// CreateBook function to create a book
func CreateBook(c *fiber.Ctx) error {
	books := datatypes.Book{}

	if err := c.BodyParser(&books); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	fmt.Println(books)
	id, error := postgres.CreateBook(books)
	if id == 0 || error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": error.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Book has been created",
		"id":      id,
	})
}

// GetBooks function to get all books
func GetBooks(c *fiber.Ctx) error {
	books, err := postgres.GetBooks()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"books": books,
	})
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	book, err := postgres.GetBook(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	error := postgres.DeleteBook(id)
	if error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": error.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Book has been deleted",
	})
}
