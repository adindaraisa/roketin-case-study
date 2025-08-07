package main

import (
	"log"
	"net/http"

	"github.com/adindaraisa/roketin-case-study/challenge-dua/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/movies", handlers.ListMovies)
	r.Post("/movies", handlers.CreateMovie)
	r.Put("/movies/{id}", handlers.UpdateMovie)
	r.Get("/movies/search", handlers.SearchMovies)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
