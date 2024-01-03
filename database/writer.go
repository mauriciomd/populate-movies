package database

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mauriciomd/populate-movies/parser"
)

type MovieWriter interface {
	WriteChunk([]parser.Movie) error
}

type BDMovieWriter struct{}

func (d BDMovieWriter) WriteChunk(chunk []parser.Movie) error {
	db := NewConnection()
	defer db.Close()

	query := "INSERT INTO populate_cli.movies (id, title, release_date, genres) VALUES "

	var values []interface{}
	mult := 1
	for i, m := range chunk {
		query += fmt.Sprintf(" ($%d, $%d, $%d, $%d)", mult+i, mult+i+1, mult+i+2, mult+i+3)
		if i < len(chunk)-1 {
			query += ", "
		}

		values = append(values, m.Id, m.Title, m.ReleaseDate, strings.Join(m.Genres, ", "))
		mult += 3
	}
	query += " ON CONFLICT (id) DO UPDATE SET created_at = NOW();"

	_, err := db.Exec(query, values...)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to insert with query=%s\n error=%s\n\n", query, err.Error()))
	}

	return nil
}
