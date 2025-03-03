package booksroutes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/rajivreddy/go-fiber-pgsql/docs"
	login "github.com/rajivreddy/go-fiber-pgsql/pkg/http/handlers/auth"
	handlers "github.com/rajivreddy/go-fiber-pgsql/pkg/http/handlers/books"
	users "github.com/rajivreddy/go-fiber-pgsql/pkg/http/handlers/users"
)

//Add Fiber Routes here

func SetupRoutes(route fiber.Router) {
	log.Println("Setting up routes")
	log.Println("Setting up routes for login")
	auth := route.Group("/auth")
	auth.Post("/login", login.Login)

	user := route.Group("/users")
	user.Post("/", users.CreateUser)
	user.Get("/:username", users.GetUser)

	log.Println("Setting up routes for Books API")
	books := route.Group("/books")
	books.Post("/", handlers.CreateBook)
	books.Get("/", handlers.GetBooks)
	books.Get("/:id", handlers.GetBook)
	// books.Put("/:id", handlers.UpdateBook)
	books.Delete("/:id", handlers.DeleteBook)
}
