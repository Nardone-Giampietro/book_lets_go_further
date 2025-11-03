package data

import (
	"time"

	"greenlight.nardone.xyz/internal/validator"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	// Call the Runtime type in order to MarshalJSON to work
	Runtime Runtime  `json:"runtime,omitempty"`
	Genres  []string `json:"genres,omitempty"`
	Version int32    `json:"version"`
}

func ValidateMovie(v *validator.Validator, inputMovie *Movie) {
	v.Check(inputMovie.Title != "", "title", "must be provided")
	v.Check(len(inputMovie.Title) <= 500, "title", "must not be more than 500 bytes")

	v.Check(inputMovie.Year != 0, "year", "must be provided")
	v.Check(inputMovie.Year >= 1888, "year", "must be greater than or equal to 1888")
	v.Check(inputMovie.Year <= int32(time.Now().Year()), "year", "must be less th current year")

	v.Check(inputMovie.Runtime != 0, "runtime", "must be provided")
	v.Check(inputMovie.Runtime > 0, "runtime", "must be greater than or equal to 0")

	v.Check(inputMovie.Genres != nil, "genres", "must be provided")
	v.Check(len(inputMovie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(inputMovie.Genres) <= 5, "genres", "must contain at least 5 genres")
	v.Check(validator.Unique(inputMovie.Genres), "genres", "must be unique")
}
