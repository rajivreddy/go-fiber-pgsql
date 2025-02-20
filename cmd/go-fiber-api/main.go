package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rajivreddy/go-fiber-pgsql/pkg/config"
	booksroutes "github.com/rajivreddy/go-fiber-pgsql/pkg/http/routes/books"
	"github.com/rajivreddy/go-fiber-pgsql/pkg/postgres"
)

func main() {

	// Load Config
	Cfg := config.LoadConfig()
	//Database connection
	err := postgres.AutoMigrate(Cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Start Fiber Server
	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api")
	booksroutes.SetupRoutes(api)
	// Start the server on port 3000

	log.Printf("Server is running on http://%s:%s", Cfg.Server.Host, Cfg.Server.Port)
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", Cfg.Server.Host, Cfg.Server.Port)))

}
