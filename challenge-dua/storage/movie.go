package storage

import (
	"strings"

	"github.com/adindaraisa/roketin-case-study/challenge-dua/models"
)

var movies []models.Movie
var nextID = 1

// AddMovie menyimpan film baru ke memori
func AddMovie(movie models.Movie) models.Movie {
	movie.ID = nextID
	nextID++
	movies = append(movies, movie)
	return movie
}

// UpdateMovie memperbarui data film
func UpdateMovie(id int, updated models.Movie) (models.Movie, bool) {
	for i, m := range movies {
		if m.ID == id {
			updated.ID = id
			movies[i] = updated
			return updated, true
		}
	}
	return models.Movie{}, false
}

// GetMovies mengembalikan list film dengan pagination
func GetMovies(offset, limit int) []models.Movie {
	if offset > len(movies) {
		return []models.Movie{}
	}
	end := offset + limit
	if end > len(movies) {
		end = len(movies)
	}
	return movies[offset:end]
}

// SearchMovies mencari film berdasarkan title, description, artists, atau genres
func SearchMovies(query string) []models.Movie {
	var result []models.Movie
	q := strings.ToLower(query)

	for _, m := range movies {
		if contains(m.Title, q) ||
			contains(m.Description, q) ||
			sliceContains(m.Artists, q) ||
			sliceContains(m.Genres, q) {
			result = append(result, m)
		}
	}

	return result
}

func contains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), substr)
}

func sliceContains(slice []string, substr string) bool {
	for _, item := range slice {
		if contains(item, substr) {
			return true
		}
	}
	return false
}
