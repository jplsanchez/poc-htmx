package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Context struct {
	db *sql.DB
}

func NewContext() (*Context, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	connStr := os.Getenv("DSN")
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping: %v", err)
	}

	log.Println("Successfully connected to PlanetScale!")

	return &Context{db}, nil
}

func (ctx *Context) Close() error {
	return ctx.db.Close()
}

func (ctx *Context) GetAllFilms() (*[]Film, error) {
	results, err := ctx.db.Query("SELECT Title, Director FROM Film")
	if err != nil {
		return nil, err
	}

	var films []Film
	for results.Next() {
		var film Film
		if err := results.Scan(&film.Title, &film.Director); err != nil {
			return nil, err
		}
		films = append(films, film)
	}
	return &films, nil
}

func (ctx *Context) InsertFilm(film Film) error {
	insert, err := ctx.db.Query(fmt.Sprintf("INSERT INTO Film VALUES ( '%v', '%v' )", film.Title, film.Director))
	if err != nil {
		return err
	}
	defer insert.Close()

	return nil
}
