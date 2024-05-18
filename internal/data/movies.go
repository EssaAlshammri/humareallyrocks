package data

import (
	"database/sql"
	"time"
)

type Movie struct {
	ID        uint      `json:"ID" doc:"Unique integer ID for the movie"`
	CreatedAt time.Time `json:"createdAt" doc:"Timestamp for when the movie is added to our database"`
	Title     string    `json:"title" doc:"Movie title"`
	Year      int32     `json:"year" doc:"Movie release year"`
	Runtime   int32     `json:"runtime" doc:"Movie runtime (in minutes)"`
	Genres    []string  `json:"genres" doc:"genres for the movie"`
}

type MovieOut struct {
	ID        uint      `json:"ID" doc:"Unique integer ID for the movie"`
	CreatedAt time.Time `json:"createdAt" doc:"Timestamp for when the movie is added to our database"`
}

type MovieIn struct {
	Title     string    `json:"title" doc:"Movie title"`
	Year      int32     `json:"year" doc:"Movie release year"`
	Runtime   int32     `json:"runtime" doc:"Movie runtime (in minutes)"`
	Genres    []string  `json:"genres" doc:"genres for the movie"`
}

type MovieModel struct {
	DB *sql.DB
}

func (m MovieModel) Insert(movie *Movie) error {
	query := `
		INSERT INTO movies (title, year, runtime, genres) 
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`

	args := []any{movie.Title, movie.Year, movie.Runtime, movie.Genres}

	return m.DB.QueryRow(query, args...).Scan(&movie.ID, &movie.CreatedAt)
}

func (m MovieModel) Get(id int64) (*Movie, error) {
	return nil, nil
}

func (m MovieModel) Update(movie *Movie) error {
	return nil
}

func (m MovieModel) Delete(id int64) error {
	return nil
}
