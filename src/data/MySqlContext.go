package data

import (
	"database/sql"
	"fmt"
	"htmx/src/model"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type MySqlContext struct {
	db *sql.DB
}

func NewMySqlContext() (*MySqlContext, error) {
	godotenv.Load()

	connStr := os.Getenv("DSN")
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping: %v", err)
	}

	return &MySqlContext{db}, nil
}

func (ctx *MySqlContext) Close() error {
	return ctx.db.Close()
}

func (ctx MySqlContext) GetAllFilms() (*[]model.Film, error) {
	results, err := ctx.db.Query("SELECT Title, Director FROM Film ORDER BY RegisterDate DESC LIMIT 10")
	if err != nil {
		return nil, err
	}

	var films []model.Film
	for results.Next() {
		var film model.Film
		if err := results.Scan(&film.Title, &film.Director); err != nil {
			return nil, err
		}
		films = append(films, film)
	}
	return &films, nil
}

func (ctx *MySqlContext) InsertFilm(film model.Film) error {
	insert, err := ctx.db.Query(fmt.Sprintf("INSERT INTO Film (Title ,Director) VALUES ( '%v', '%v' )", film.Title, film.Director))
	if err != nil {
		return err
	}
	defer insert.Close()

	return nil
}
