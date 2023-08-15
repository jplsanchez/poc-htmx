package data

import (
	"htmx/src/model"

	_ "github.com/go-sql-driver/mysql"
)

type Repository interface {
	GetAllFilms() (*[]model.Film, error)
	InsertFilm(film model.Film) error
	Close() error
}
