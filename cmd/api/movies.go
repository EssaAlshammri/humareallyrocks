package main

import (
	"context"
	"time"

	"github.com/EssaAlshammri/humareallyrocks/internal/data"
)

func (app application) createMovieHandler(ctx context.Context, input *MovieCreateInput) (*MovieCreateOut, error) {
	movie := &data.Movie{
		Title:   input.Body.Title,
		Year:    input.Body.Year,
		Runtime: input.Body.Runtime,
		Genres:  input.Body.Genres,
	}
	err := app.models.Movies.Insert(movie)
	if err != nil {
		return nil, err
	}
	output := &MovieCreateOut{}
	output.Body.ID = movie.ID
	output.Body.CreatedAt = movie.CreatedAt
	return output, nil
}

func (app application) getMovieHandler(ctx context.Context, input *MovieGetInput) (*MovieGetOutput, error) {
	id := input.ID
	resp := &MovieGetOutput{}
	resp.Body.ID = id
	resp.Body.Title = "hello"
	resp.Body.CreatedAt = time.Now().UTC()
	return resp, nil
}
