package booksroutes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/rajivreddy/go-fiber-pgsql/docs"
	handlers "github.com/rajivreddy/go-fiber-pgsql/pkg/http/handlers/books"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

//Add Fiber Routes here

func SetupRoutes(route fiber.Router) {
	log.Println("Setting up routes")
	route.Get("/swagger/*", fiberSwagger.WrapHandler)
	books := route.Group("/books")
	books.Post("/", handlers.CreateBook)
	books.Get("/", handlers.GetBooks)
	books.Get("/:id", handlers.GetBook)
	// books.Put("/:id", handlers.UpdateBook)
	books.Delete("/:id", handlers.DeleteBook)
}
