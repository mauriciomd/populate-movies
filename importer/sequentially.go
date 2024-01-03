package importer

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/mauriciomd/populate-movies/database"
	"github.com/mauriciomd/populate-movies/errors"
	"github.com/mauriciomd/populate-movies/parser"
)

type Sequentially struct {
	File      *os.File
	Writer    database.BDMovieWriter
	ChunkSize int
}

func (s *Sequentially) Process() {
	errors := errors.New()

	reader := csv.NewReader(s.File)
	dataChunk := []parser.Movie{}
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
		if len(dataChunk)%s.ChunkSize == 0 {
			dataChunk = s.writeData(dataChunk, errors)
		}
	}

	if len(dataChunk) > 0 {
		dataChunk = s.writeData(dataChunk, errors)
	}

	if errors.Len() > 0 {
		log.Fatalln(errors.Error())
	}
}

func (s Sequentially) writeData(dataChunk []parser.Movie, e *errors.ErrorList) []parser.Movie {
	err := s.Writer.WriteChunk(dataChunk)
	if err != nil {
		e.Add(err)
	}

	return []parser.Movie{}
}
