package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	ID     = 0
	TITLE  = 1
	GENRES = 2
)

func GetMovie(line []string) (Movie, error) {
	if len(line) < 3 {
		return Movie{}, errors.New("Invalid line.")
	}

	y, err := getYearFromTitle(line[TITLE])
	if err != nil {
		return Movie{}, err
	}

	year, err := strconv.Atoi(y)
	if err != nil {
		return Movie{}, err
	}

	title := removeYearFromTitle(clean(line[TITLE]), year)
	genres := separateGenres(line[GENRES])

	id, err := strconv.Atoi(line[ID])

	return NewMovie(title, genres, id, year), nil
}

func clean(s string) string {
	l := strings.ReplaceAll(s, "\"", "")
	return strings.ReplaceAll(l, "\n", "")
}

func getYearFromTitle(str string) (string, error) {
	r, err := regexp.Compile("[0-9]{4}")
	if err != nil {
		return "", err
	}

	idx := r.FindStringIndex(str)

	// There's no year in the title
	if len(idx) == 0 {
		return "0", nil
	}

	s, e := idx[0], idx[1]
	return string(str[s:e]), nil
}

func removeYearFromTitle(s string, year int) string {
	sYear := fmt.Sprintf("(%d)", year)
	return strings.ReplaceAll(s, sYear, "")
}

func separateGenres(s string) []string {
	return strings.Split(s, "|")
}

// func split(s string) []string {
// 	runes := []rune(s)
// 	firstComma := strings.IndexRune(s, ',')
// 	lastComma := strings.LastIndexAny(s, ",")

// 	id := string(runes[:firstComma])
// 	title := string(runes[firstComma+1 : lastComma])
// 	genres := string(runes[lastComma+1:])

// 	return []string{id, title, genres}
// }
