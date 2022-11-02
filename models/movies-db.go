package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModels struct {
	DB *sql.DB
}

// Get return one movie and error, if any.
func (m *DBModels) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, description, year, release_date, rating, runtime, mpaa_rating,
				created_at, updated_at from movies where id = $1
	`

	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie
	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Rating,
		&movie.Runtime,
		&movie.MPAARating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

// All return all movie and error, if any.
func (m *DBModels) All() ([]*Movie, error) {
	return nil, nil
}
