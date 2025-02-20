package postgres

import (
	"fmt"
	"log"

	"github.com/rajivreddy/go-fiber-pgsql/pkg/config"
	"github.com/rajivreddy/go-fiber-pgsql/pkg/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(cfg *config.Config) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Db.Host, cfg.Db.Username, cfg.Db.Password, cfg.Db.DbName, cfg.Db.Port)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func AutoMigrate(cfg *config.Config) error {
	if db == nil {
		if err := InitDB(cfg); err != nil {
			return err
		}
	}
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

func CreateBook(book datatypes.Book) (int, error) {
	if db == nil {
		return 0, fmt.Errorf("database not initialized")
	}
	err := db.Create(&book).Error
	if err != nil {
		return 0, err
	}
	log.Printf("Book has been created with id %d", book.ID)
	return book.ID, nil
}

func GetBooks() ([]datatypes.Book, error) {
	if db == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	var books []datatypes.Book
	err := db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func GetBook(id string) (datatypes.Book, error) {
	if db == nil {
		return datatypes.Book{}, fmt.Errorf("database not initialized")
	}
	var book datatypes.Book
	err := db.First(&book, id).Error
	if err != nil {
		return datatypes.Book{}, err
	}
	return book, nil
}

func DeleteBook(id string) error {
	if db == nil {
		return fmt.Errorf("database not initialized")
	}
	err := db.Delete(&datatypes.Book{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
