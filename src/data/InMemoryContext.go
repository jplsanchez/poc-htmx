package data

import (
	"htmx/src/model"
)

type InMemoryContext struct {
	list []model.Film
}

func NewInMemoryContext() (*InMemoryContext, error) {

	list := []model.Film{
		{Title: "The Shawshank Redemption", Director: "Frank Darabont"},
		{Title: "The Godfather", Director: "Francis Ford Coppola"},
		{Title: "The Godfather: Part II", Director: "Francis Ford Coppola"},
	}

	return &InMemoryContext{list}, nil
}

func (ctx *InMemoryContext) Close() error {
	ctx.list = make([]model.Film, 0)
	return nil
}

func (ctx *InMemoryContext) GetAllFilms() (*[]model.Film, error) {
	return &ctx.list, nil
}

func (ctx *InMemoryContext) InsertFilm(film model.Film) error {
	ctx.list = append(ctx.list, film)
	return nil
}
