package database

type Database interface {
	Create(name string, author string) error
}
