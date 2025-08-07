package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/adindaraisa/roketin-case-study/challenge-dua-with-database/config"
	"github.com/adindaraisa/roketin-case-study/challenge-dua-with-database/models"
	"github.com/adindaraisa/roketin-case-study/challenge-dua-with-database/repositories"
	"github.com/gin-gonic/gin"
)

type CreateMovieRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Duration    int      `json:"duration"`
	Artists     []string `json:"artists"`
	Genres      []string `json:"genres"`
}

func CreateMovie(c *gin.Context) {
	var input CreateMovieRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie := models.Movie{
		Title:       input.Title,
		Description: input.Description,
		Duration:    input.Duration,
	}
	config.DB.Create(&movie)

	for _, name := range input.Artists {
		var artist models.Artist
		config.DB.FirstOrCreate(&artist, models.Artist{Name: name})
		config.DB.Model(&movie).Association("Artists").Append(&artist)
	}

	for _, name := range input.Genres {
		var genre models.Genre
		config.DB.FirstOrCreate(&genre, models.Genre{Name: name})
		config.DB.Model(&movie).Association("Genres").Append(&genre)
	}

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

func UpdateMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	var req CreateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var movie models.Movie
	if err := config.DB.Preload("Artists").Preload("Genres").First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	movie.Title = req.Title
	movie.Description = req.Description
	movie.Duration = req.Duration

	config.DB.Model(&movie).Association("Artists").Clear()
	config.DB.Model(&movie).Association("Genres").Clear()

	for _, name := range req.Artists {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}

		var artist models.Artist
		if err := config.DB.FirstOrCreate(&artist, models.Artist{Name: name}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create/find artist"})
			return
		}

		config.DB.Model(&movie).Association("Artists").Append(&artist)
	}

	for _, name := range req.Genres {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}

		var genre models.Genre
		if err := config.DB.FirstOrCreate(&genre, models.Genre{Name: name}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create/find genre"})
			return
		}

		config.DB.Model(&movie).Association("Genres").Append(&genre)
	}

	if err := config.DB.Save(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

func ListMovies(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	repo := repositories.NewMovieRepository(config.DB)
	movies, err := repo.FindAll(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, movies)
}

func SearchMovies(c *gin.Context) {
	query := c.Query("q")

	repo := repositories.NewMovieRepository(config.DB)
	movies, err := repo.Search(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, movies)
}
