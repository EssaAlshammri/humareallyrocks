package main

import (
	"context"
	"time"
)

func (app application) createMovieHandler(ctx context.Context, input *MovieCreateInput) (*struct{}, error) {
	return nil, nil
}

func (app application) getMovieHandler(ctx context.Context, input *MovieGetInput) (*MovieGetOutput, error) {
	id := input.ID
	resp := &MovieGetOutput{}
	resp.Body.ID = id
	resp.Body.Title = "hello"
	resp.Body.CreatedAt = time.Now().UTC()
	return resp, nil
}
