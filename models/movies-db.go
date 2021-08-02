package models

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	Query := `SELECT id, title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at
   FROM public.movies where id=$1`
	row := m.DB.QueryRowContext(ctx, Query, id)
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
		log.Println("Error occor in query")
		return nil, err
	}
	Query = `select mg.id,mg.movie_id,mg.genre_id,g.genre_name from
		   movies_genres mg left join genres g on (g.id=mg.genre_id)
		   where 
		   mg.movie_id=$1
	`
	rows, _ := m.DB.QueryContext(ctx, Query, id)
	defer rows.Close()

	var genres []MovieGenre

	for rows.Next() {
		var mg MovieGenre
		err := rows.Scan(
			&mg.ID,
			&mg.MovieID,
			&mg.GenreID,
			&mg.Genre.GenreName,
		)
		if err != nil {
			log.Println("custom error is ", err)
			return nil, err
		}
		genres = append(genres, mg)
	}
	movie.MovieGenre = genres

	return &movie, nil

}
func (m *DBModel) All() ([]*Movie, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	Query := `SELECT id, title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at
  FROM public.movies`
	rows, err := m.DB.QueryContext(ctx, Query)
	if err != nil {
		return nil, err
	}
	var movies []*Movie
	for rows.Next() {
		var movie Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.Rating,
			&movie.MPAARating,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		var getMovies []MovieGenre
		genreQuery := `select mg.id,mg.movie_id,mg.genre_id,g.genre_name from
						movies_genres mg left join genres g on (g.id=mg.genre_id)
						where 
						mg.movie_id=$1
                   `
		genreRows, _ := m.DB.QueryContext(ctx, genreQuery, movie.ID)
		for genreRows.Next() {
			var mg MovieGenre
			err := genreRows.Scan(
				&mg.ID,
				&mg.MovieID,
				&mg.GenreID,
				&mg.Genre.GenreName,
			)
			if err != nil {
				return nil, err
			}
			getMovies = append(getMovies, mg)
		}
		genreRows.Close()
		movie.MovieGenre = getMovies
		movies = append(movies, &movie)
	}

	return movies, nil
}
