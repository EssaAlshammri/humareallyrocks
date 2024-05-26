package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Movie struct {
	ID        uint      `json:"ID" doc:"Unique integer ID for the movie" minimum:"1"`
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
	Title   string   `json:"title" doc:"Movie title"`
	Year    int32    `json:"year" doc:"Movie release year" minimum:"1888"`
	Runtime int32    `json:"runtime" doc:"Movie runtime (in minutes)"`
	Genres  []string `json:"genres" doc:"genres for the movie"`
}

type MovieModel struct {
	DB *sql.DB
}

func (m MovieModel) Insert(ctx context.Context, movieIn *MovieIn) (*MovieOut, error) {
	dbCtx, cancel := context.WithTimeout(ctx, WriteTimeout)
	defer cancel()
	query := `
		INSERT INTO movies (title, year, runtime, genres) 
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`

	args := []any{movieIn.Title, movieIn.Year, movieIn.Runtime, movieIn.Genres}

	var movieOut MovieOut
	err := m.DB.QueryRowContext(dbCtx, query, args...).Scan(&movieOut.ID, &movieOut.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &movieOut, nil
}

func (m MovieModel) Get(ctx context.Context, id int64) (*Movie, error) {
	dbCtx, cancel := context.WithTimeout(ctx, ReadTimeout)
	defer cancel()
	query := `
		SELECT id, created_at, title, year, runtime, genres
		FROM movies
		WHERE id = $1
	`

	var movie Movie

	mapper := pgtype.NewMap()

	err := m.DB.QueryRowContext(dbCtx, query, id).Scan(
		&movie.ID,
		&movie.CreatedAt,
		&movie.Title,
		&movie.Year,
		&movie.Runtime,
		mapper.SQLScanner(&movie.Genres),
	)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (m MovieModel) Update(ctx context.Context, id int64, movieIn *MovieIn) (*MovieOut, error) {
	dbCtx, cancel := context.WithTimeout(ctx, WriteTimeout)
	defer cancel()
	query := `
        UPDATE movies 
        SET title = $1, year = $2, runtime = $3, genres = $4
        WHERE id = $5
		RETURNING id, created_at
	`

	args := []any{
		movieIn.Title,
		movieIn.Year,
		movieIn.Runtime,
		movieIn.Genres,
		id,
	}

	var movieOut MovieOut
	err := m.DB.QueryRowContext(dbCtx, query, args...).Scan(&movieOut.ID, &movieOut.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &movieOut, nil
}

func (m MovieModel) Delete(ctx context.Context, id int64) error {
	dbCtx, cancel := context.WithTimeout(ctx, WriteTimeout)
	defer cancel()
	query := `
        DELETE FROM movies
        WHERE id = $1
	`

	result, err := m.DB.ExecContext(dbCtx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
