package main

import (
	"context"

	"github.com/EssaAlshammri/humareallyrocks/internal/data"
	"github.com/danielgtaylor/huma/v2"
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
		return nil, huma.Error404NotFound("movie not found")
	}

	output := &MovieGetOut{}
	output.Body = struct{ data.Movie }{*movie}
	return output, nil
}

func (app application) updateMovieHandler(ctx context.Context, input *MovieUpdateIn) (*MovieUpdateOut, error) {
	_, err := app.models.Movies.Get(input.ID)
	if err != nil {
		return nil, huma.Error404NotFound("movie not found")
	}

	movieOut, err := app.models.Movies.Update(input.ID, &input.Body.MovieIn)
	if err != nil {
		return nil, err
	}
	output := &MovieUpdateOut{}
	output.Body = struct{ data.MovieOut }{*movieOut}
	return output, nil
}
