package repositories

import (
	"gorm.io/gorm"

	"github.com/adindaraisa/roketin-case-study/challenge-dua-with-database/models"
)

type MovieRepository struct {
	DB *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{DB: db}
}

func (r *MovieRepository) Create(movie *models.Movie) error {
	return r.DB.Create(movie).Error
}

func (r *MovieRepository) Update(movie *models.Movie) error {
	return r.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(movie).Error
}

func (r *MovieRepository) FindAll(offset, limit int) ([]models.Movie, error) {
	var movies []models.Movie
	err := r.DB.Preload("Artists").Preload("Genres").Offset(offset).Limit(limit).Find(&movies).Error
	return movies, err
}

func (r *MovieRepository) Search(query string) ([]models.Movie, error) {
	var movies []models.Movie
	err := r.DB.
		Preload("Artists").
		Preload("Genres").
		Joins("LEFT JOIN movie_artists ON movie_artists.movie_id = movies.id").
		Joins("LEFT JOIN artists ON artists.id = movie_artists.artist_id").
		Joins("LEFT JOIN movie_genres ON movie_genres.movie_id = movies.id").
		Joins("LEFT JOIN genres ON genres.id = movie_genres.genre_id").
		Where("movies.title ILIKE ? OR movies.description ILIKE ? OR artists.name ILIKE ? OR genres.name ILIKE ?",
			"%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%").
		Group("movies.id").
		Find(&movies).Error
	return movies, err
}
