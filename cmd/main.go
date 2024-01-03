package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mauriciomd/populate-movies/database"
	"github.com/mauriciomd/populate-movies/importer"
)

func init() {
	log.SetOutput(os.Stdout)
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func parseArgs() (string, int, bool, int) {
	file := flag.String("f", "data.csv", "File containing all movies (.csv comma separated).")
	chunkSize := flag.Int("s", 8, "Chunk size used to insert on database.")
	isConcurrent := flag.Bool("c", false, "Set CLI to operate on concurrent mode.")
	workers := flag.Int("w", 4, "Set the number of goroutines to import the dataset.")
	flag.Parse()

	return *file, *chunkSize, *isConcurrent, *workers
}

func formatLog(filename string, chunksize, workers int, isConcurrent bool) string {
	var str string
	if isConcurrent {
		str = fmt.Sprintf("Running CLI on concurrentlty mode with workers=%d ", workers)
	} else {
		str = "Running CLI on sequentially mode "
	}

	str += fmt.Sprintf("and chunkSize=%d with file=%s", chunksize, filename)

	return str
}

func main() {
	filename, chunkSize, isConcurrent, workers := parseArgs()
	message := formatLog(filename, chunkSize, workers, isConcurrent)
	log.Println(message)

	file, err := os.Open(filename)
	checkError(err)
	defer file.Close()

	var worker importer.Worker
	worker = &importer.Sequentially{
		File:      file,
		Writer:    database.BDMovieWriter{},
		ChunkSize: chunkSize,
	}

	if isConcurrent {
		worker = &importer.Concurrently{
			File:      file,
			Writer:    database.BDMovieWriter{},
			ChunkSize: chunkSize,
			Workers:   workers,
		}
	}

	worker.Process()
}
