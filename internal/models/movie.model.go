package models

import (
	"time"

	"gorm.io/datatypes"
)

type Movie struct {
	ID              uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Title           string         `gorm:"type:varchar(255);not null" json:"title" validate:"required"`
	Description     string         `gorm:"type:text;not null" json:"description" validate:"required"`
	PosterURL       string         `gorm:"type:varchar(255);not null" json:"poster_url" validate:"required,url"`
	ReleaseDate     string         `gorm:"type:date;not null" json:"release_date" validate:"required,datetime=2006-01-02"`
	Rating          float64        `gorm:"type:decimal(3,1);not null" json:"rating" validate:"required,numeric"`
	DurationMinutes int            `gorm:"type:int;not null" json:"duration_minutes" validate:"required,numeric"`
	Director        string         `gorm:"type:varchar(255);not null" json:"director" validate:"required"`
	Genre           datatypes.JSON `gorm:"type:json;not null" json:"genre" validate:"required,min=1,dive"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}
