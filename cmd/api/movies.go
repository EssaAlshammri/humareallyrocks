package main

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
)

func (app application) createMovieHandler(ctx context.Context, input *MovieCreateIn) (*MovieCreateOut, error) {
	movieOut, err := app.models.Movies.Insert(ctx, &input.Body)
	if err != nil {
		return nil, err
	}
	output := &MovieCreateOut{}
	output.Body = *movieOut
	return output, nil
}

func (app application) getMovieHandler(ctx context.Context, input *MovieGetIn) (*MovieGetOut, error) {
	movie, err := app.models.Movies.Get(ctx, input.ID)
	if err != nil {
		return nil, huma.Error404NotFound("movie not found")
	}

	output := &MovieGetOut{}
	output.Body = *movie
	return output, nil
}

func (app application) updateMovieHandler(ctx context.Context, input *MovieUpdateIn) (*MovieUpdateOut, error) {
	_, err := app.models.Movies.Get(ctx, input.ID)
	if err != nil {
		return nil, huma.Error404NotFound("movie not found")
	}

	movieOut, err := app.models.Movies.Update(ctx, input.ID, &input.Body)
	if err != nil {
		return nil, err
	}
	output := &MovieUpdateOut{}
	output.Body = *movieOut
	return output, nil
}

func (app application) deleteMovieHandler(ctx context.Context, input *MovieDeleteIn) (*struct{}, error) {
	_, err := app.models.Movies.Get(ctx, input.ID)
	if err != nil {
		return nil, huma.Error404NotFound("movie not found")
	}

	err = app.models.Movies.Delete(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (app application) listMoviesHandler(ctx context.Context, input *MoviesListIn) (*MoviesListOut, error) {
	movies, err := app.models.Movies.GetAll(ctx, input.Title, input.Sort, input.Genres, input.Page, input.PageSize)
	if err != nil {
		app.logger.Info(err.Error())
		return nil, huma.Error400BadRequest("bad request")
	}
	output := &MoviesListOut{}
	output.Body.Movies = *movies
	return output, nil
}
