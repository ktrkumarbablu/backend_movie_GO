package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameters"))
		app.errorJson(w, err)
		return
	}
	app.logger.Println("id is", id)
	movie, err := app.models.DB.Get(id)
	if err != nil {
		app.logger.Println("error")
	}

	// movie := models.Movie{
	// 	ID:          id,
	// 	Title:       "Some Movie",
	// 	Description: "Some description",
	// 	Year:        2021,
	// 	ReleaseDate: time.Date(2021, 01, 01, 01, 0, 0, 0, time.Local),
	// 	Runtime:     100,
	// 	Rating:      5,
	// 	MPAARating:  "PG-13",
	// 	CreatedAt:   time.Now(),
	// 	UpdatedAt:   time.Now(),
	// }
	err = app.writeJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.logger.Println("error")
	}
}
func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.All()
	if err != nil {
		app.errorJson(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJson(w, err)
	}

}
func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request)        {}
func (app *application) deleteMovieByID(w http.ResponseWriter, r *http.Request)    {}
func (app *application) insertNewMovie(w http.ResponseWriter, r *http.Request)     {}
func (app *application) UpdateNewMovieByID(w http.ResponseWriter, r *http.Request) {}
