package main

import (
	"context"

	"github.com/EssaAlshammri/humareallyrocks/internal/data"
)

func (app application) createMovieHandler(ctx context.Context, input *MovieCreateIn) (*MovieCreateOut, error) {
	movieOut, err := app.models.Movies.Insert(&input.Body.MovieIn)
	if err != nil {
		return nil, err
	}
	output := &MovieCreateOut{}
	output.Body = struct{ data.MovieOut }{*movieOut}
	return output, nil
}

func (app application) getMovieHandler(ctx context.Context, input *MovieGetIn) (*MovieGetOut, error) {
	movie, err := app.models.Movies.Get(input.ID)
	if err != nil {
		return nil, err
	}

	output := &MovieGetOut{}
	output.Body = struct{ data.Movie }{*movie}
	return output, nil
}
