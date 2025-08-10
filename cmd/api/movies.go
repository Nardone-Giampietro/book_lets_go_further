package main

import (
	"fmt"
	"greenlight.nardone.xyz/internal/data"
	"net/http"
	"time"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	// This the object that we expect to receive from the request
	var inputMovie struct {
		Title   string   `json:"title"`
		Year    int      `json:"year"`
		Runtime int32    `json:"runtime"`
		Genre   []string `json:"genres"`
	}

	err := app.readJSON(w, r, &inputMovie)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	fmt.Fprintf(w, "%+v", inputMovie)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   2,
		Genres:    []string{"drama", "romance"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
