package postgres

import (
	"fmt"
	"log"

	"github.com/rajivreddy/go-fiber-pgsql/pkg/config"
	"github.com/rajivreddy/go-fiber-pgsql/pkg/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func AutoMigrate(cfg *config.Config) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Db.Host, cfg.Db.Username, cfg.Db.Password, cfg.Db.DbName, cfg.Db.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	error := db.AutoMigrate(&datatypes.Book{})
	if error != nil {
		log.Fatal(error)
		return error
	}
	log.Printf("Auto Migration has been completed")
	return nil
}

// func NewDB(cfg config.Config) (*gorm.DB, error) {
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Db.Host, cfg.Db.Username, cfg.Db.Password, cfg.Db.DbName, cfg.Db.Port)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	db.Exec(`
// 	CREATE TABLE IF NOT EXISTS books (
//     ID SERIAL PRIMARY KEY,
//     Title TEXT UNIQUE NOT NULL,
//     Author TEXT NOT NULL,
//     Publisher TEXT NOT NULL,
//     YearOfPublication INTEGER NOT NULL
// 	);`)

// 	return db, nil
// }

// func
