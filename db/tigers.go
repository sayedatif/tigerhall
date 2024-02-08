package db

import (
	"time"

	"gorm.io/gorm"
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

func (t *Tiger) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	t.CreatedAt = now
	t.UpdatedAt = now
	return nil
}

func (t *Tiger) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = time.Now()
	return nil
}
