package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/adindaraisa/roketin-case-study/challenge-dua/models"
	"github.com/adindaraisa/roketin-case-study/challenge-dua/storage"
	"github.com/go-chi/chi/v5"
)

// Helper untuk output JSON yang rapi
func writePrettyJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	out, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(out)
}

// CreateMovie POST /movies
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	created := storage.AddMovie(movie)
	writePrettyJSON(w, created)
}

// UpdateMovie PUT /movies/{id}
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, ok := storage.UpdateMovie(id, movie)
	if !ok {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	writePrettyJSON(w, updated)
}

// ListMovies GET /movies?page=1&limit=10
func ListMovies(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	movies := storage.GetMovies(offset, limit)
	writePrettyJSON(w, movies)
}

// SearchMovies GET /movies/search?query=...
func SearchMovies(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("query")))
	if query == "" {
		http.Error(w, "Query parameter is required", http.StatusBadRequest)
		return
	}

	results := storage.SearchMovies(query)
	writePrettyJSON(w, results)
}
