package db

import (
	"time"
)

type Tiger struct {
	ID           int64  `gorm:"primaryKey;autoIncrement;not null"`
	Name         string `gorm:"type:varchar(255);not null"`
	DOB          string
	LastSeenAt   time.Time
	LastSeenLat  float64
	LastSeenLong float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Users        []User `gorm:"many2many:user_tiger_sightings"`
}
