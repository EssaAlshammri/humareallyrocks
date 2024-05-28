package main

import (
	"reflect"

	"github.com/EssaAlshammri/humareallyrocks/internal/data"
	"github.com/danielgtaylor/huma/v2"
)

type HealthCheckOutput struct {
	Body struct {
		Status      string `json:"status" example:"available" doc:"server status"`
		Environment string `json:"environment" example:"development|staging|production" doc:"running environment"`
		Version     string `json:"version" example:"1.0.0" doc:"application version"`
	}
}

type PaginationParams struct {
	Page     int64 `query:"page" minimum:"1" default:"1"`
	PageSize int64 `query:"page_size" minimum:"1" maximum:"100" default:"10"`
}

type MovieCreateIn struct {
	Body data.MovieIn
}

type MovieCreateOut struct {
	Body data.MovieOut
}

type MovieGetIn struct {
	ID int64 `path:"id" minimum:"1"`
}

type MovieGetOut struct {
	Body data.Movie
}

type MovieUpdateIn struct {
	ID   int64 `path:"id" minimum:"1"`
	Body data.MovieIn
}

type MovieUpdateOut struct {
	Body data.MovieOut
}

type MovieDeleteIn struct {
	ID int64 `path:"id" minimum:"1"`
}

type MoviesListIn struct {
	Title  string   `query:"title" doc:"title of the movie"`
	Genres []string `query:"genres" doc:"genres of the movie"`
	Sort   string   `query:"sort" default:"id" enum:"id,title,year,runtime,-id,-title,-year,-runtime"`
	PaginationParams
}

func (m *MoviesListIn) Resolve(ctx huma.Context) []error {

	if reflect.ValueOf(m.Genres).IsZero() {
		m.Genres = []string{}
	}

	return nil
}

type MoviesListOut struct {
	Body struct {
		Movies []data.Movie `json:"movies"`
	}
}
