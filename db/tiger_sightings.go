package db

import (
	"time"
)

type TigerSighting struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;not null"`
	TigerId   uint      `gorm:"not null;index"`
	Lat       float64   `gorm:"not null"`
	Long      float64   `gorm:"not null"`
	SeenAt    time.Time `gorm:"not null"`
	ImageUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
