package booksroutes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	handlers "github.com/rajivreddy/go-fiber-pgsql/pkg/http/handlers/books"
)

//Add Fiber Routes here

func SetupRoutes(route fiber.Router) {
	log.Println("Setting up routes")
	books := route.Group("/books")
	books.Post("/", handlers.CreateBook)
	books.Get("/", handlers.GetBooks)
	books.Get("/:id", handlers.GetBook)
	// books.Put("/:id", handlers.UpdateBook)
	books.Delete("/:id", handlers.DeleteBook)
}
