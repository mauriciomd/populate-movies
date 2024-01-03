package importer

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"sync"

	"github.com/mauriciomd/populate-movies/database"
	"github.com/mauriciomd/populate-movies/errors"
	"github.com/mauriciomd/populate-movies/parser"
)

type Concurrently struct {
	File      *os.File
	Writer    database.BDMovieWriter
	ChunkSize int
	Workers   int
}

func (c *Concurrently) Process() {
	errors := errors.New()

	chMovies := make(chan []parser.Movie)
	var wg sync.WaitGroup

	wg.Add(c.Workers)

	reader := csv.NewReader(c.File)
	dataChunk := []parser.Movie{}

	for i := 0; i < c.Workers; i++ {
		go c.run(chMovies, &wg, errors)
	}

	line := -1
	for {
		record, err := reader.Read()
		line++

		if err == io.EOF {
			break
		}

		if line == 0 {
			continue
		}

		movie, err := parser.GetMovie(record)
		if err != nil {
			errors.Add(err)
			continue
		}

		dataChunk = append(dataChunk, movie)
		if len(dataChunk)%c.ChunkSize == 0 {
			chMovies <- dataChunk
			dataChunk = []parser.Movie{}
		}
	}

	if len(dataChunk) > 0 {
		chMovies <- dataChunk
		dataChunk = []parser.Movie{}
	}

	close(chMovies)
	wg.Wait()

	if errors.Len() > 0 {
		log.Fatalln(errors.Error())
	}
}

func (c *Concurrently) run(chMovies chan []parser.Movie, wg *sync.WaitGroup, e *errors.ErrorList) {
	for {
		select {
		case movies, ok := <-chMovies:
			if !ok {
				wg.Done()
				return
			}

			err := c.Writer.WriteChunk(movies)
			if err != nil {
				e.Add(err)
			}
		}
	}
}
