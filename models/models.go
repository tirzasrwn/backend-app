package models

import (
	"database/sql"
	"time"
)

// Models is the wrapper for database.
type Models struct {
	DB DBModels
}

// NewModels return models with db pool.
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModels{DB: db},
	}
}

// Movie is the type of movie.
type Movie struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Year        int            `json:"year"`
	ReleaseDate time.Time      `json:"release_date"`
	Runtime     int            `json:"runtime"`
	Rating      int            `json:"rating"`
	MPAARating  string         `json:"mpaa_rating"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	MovieGenre  map[int]string `json:"genres"`
}

// Genre is the type of genre.
type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// MovieGenre is the type for movie genre.
type MovieGenre struct {
	ID        int       `json:"-"`
	MovieID   int       `json:"-"`
	GenreID   int       `json:"-"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// User is the type for users
type User struct {
	ID       int
	Email    string
	Password string
}
