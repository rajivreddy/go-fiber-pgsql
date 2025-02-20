package main

import (
	"log"

	"github.com/rajivreddy/go-fiber-pgsql/pkg/config"
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
}
