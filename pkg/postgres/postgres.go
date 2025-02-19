package postgres

import (
	"fmt"

	"github.com/rajivreddy/go-fiber-pgsql/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Db.Host, cfg.Db.Username, cfg.Db.Password, cfg.Db.DbName, cfg.Db.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.Exec(`
	CREATE TABLE IF NOT EXISTS books (
    ID SERIAL PRIMARY KEY,
    Title TEXT UNIQUE NOT NULL,
    Author TEXT NOT NULL,
    Publisher TEXT NOT NULL,
    YearOfPublication INTEGER NOT NULL
	);`)

	return db, nil
}
