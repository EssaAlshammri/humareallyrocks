package data

import "time"

type Movie struct {
	ID        uint      `json:"ID" doc:"Unique integer ID for the movie"`
	CreatedAt time.Time `json:"createdAt" doc:"Timestamp for when the movie is added to our database"`
	Title     string    `json:"title" doc:"Movie title"`
	Year      int32     `json:"year" doc:"Movie release year"`
	Runtime   int32     `json:"runtime" doc:"Movie runtime (in minutes)"`
	Genres    []string  `json:"genres" doc:"genres for the movie"`
}
