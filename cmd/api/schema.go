package main

import "github.com/EssaAlshammri/humareallyrocks/internal/data"

type HealthCheckOutput struct {
	Body struct {
		Status      string `json:"status" example:"available" doc:"server status"`
		Environment string `json:"environment" example:"development|staging|production" doc:"running environment"`
		Version     string `json:"version" example:"1.0.0" doc:"application version"`
	}
}

type MovieCreateIn struct {
	Body struct {
		data.MovieIn
	}
}

type MovieCreateOut struct {
	Body struct {
		data.MovieOut
	}
}

type MovieGetIn struct {
	ID int64 `path:"id" minimum:"1"`
}

type MovieGetOut struct {
	Body struct {
		data.Movie
	}
}

type MovieUpdateIn struct {
	ID   int64 `path:"id" minimum:"1"`
	Body struct {
		data.MovieIn
	}
}

type MovieUpdateOut struct {
	Body struct {
		data.MovieOut
	}
}


type MovieDeleteIn struct {
	ID int64 `path:"id" minimum:"1"`
}