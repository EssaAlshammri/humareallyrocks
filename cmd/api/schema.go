package main

import "time"

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
		ID        uint     `json:"ID" doc:"Unique integer ID for the movie"`
		CreatedAt time.Time `json:"createdAt" doc:"Timestamp for when the movie is added to our database"`
		Title     string    `json:"title" doc:"Movie title"`
		Year      int32     `json:"year" doc:"Movie release year"`
		Runtime   int32     `json:"runtime" doc:"Movie runtime (in minutes)"`
		Genres    []string  `json:"genres" doc:"genres for the movie"`
	}
}
