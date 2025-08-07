package models

type Artist struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
