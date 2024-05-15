package main

import "github.com/EssaAlshammri/humareallyrocks/internal/data"

type HealthCheckOutput struct {
	Body struct {
		Status      string `json:"status" example:"available" doc:"server status"`
		Environment string `json:"environment" example:"development|staging|production" doc:"running environment"`
		Version     string `json:"version" example:"1.0.0" doc:"application version"`
	}
}

type MovieCreateInput struct {
	Body struct {
		Name string `json:"name" maxLength:"50" doc:"movie name"`
	}
}

type MovieGetInput struct {
	ID uint `path:"id" minimum:"1"`
}

type MovieGetOutput struct {
	Body struct {
		data.Movie
	}
}
