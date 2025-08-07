package main

import (
	"github.com/adindaraisa/roketin-case-study/challenge-dua-with-database/config"
	"github.com/adindaraisa/roketin-case-study/challenge-dua-with-database/handlers"
	"github.com/adindaraisa/roketin-case-study/challenge-dua-with-database/models"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Movie{}, &models.Artist{}, &models.Genre{})

	r := gin.Default()

	r.POST("/movies", handlers.CreateMovie)
	r.PUT("/movies/:id", handlers.UpdateMovie)
	r.GET("/movies", handlers.ListMovies)
	r.GET("/movies/search", handlers.SearchMovies)

	r.Run(":8080")
}
