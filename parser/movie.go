package parser

import (
	"fmt"
	"time"
)

type Movie struct {
	Id          int
	Title       string
	ReleaseDate int
	Genres      []string
	CreatedAt   time.Time
}

func (m Movie) String() string {
	return fmt.Sprintf("ID: %d - Title: %s - Date: %d - Genres: %s", m.Id, m.Title, m.ReleaseDate, m.Genres)
}

func NewMovie(title string, genres []string, id, date int) Movie {
	return Movie{
		Id:          id,
		Title:       title,
		ReleaseDate: date,
		Genres:      genres,
		CreatedAt:   time.Now(),
	}
}
