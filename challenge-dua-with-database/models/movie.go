package models

import "time"

type Movie struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"` // in minutes
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Artists []Artist `gorm:"many2many:movie_artists" json:"artists"`
	Genres  []Genre  `gorm:"many2many:movie_genres" json:"genres"`
}
