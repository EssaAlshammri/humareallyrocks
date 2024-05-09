package main

import (
	"context"
)

func (app application) createMovieHandler(ctx context.Context, input *MovieCreateInput) (*struct{}, error) {
	return nil, nil
}

func (app application) getMovieHandler(ctx context.Context, input *MovieGetInput) (*MovieGetOutput, error) {
	// id := input.ID
	resp := &MovieGetOutput{}
	return resp, nil
}
