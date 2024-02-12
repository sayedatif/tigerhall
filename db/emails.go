package db

import (
	"time"

	"gorm.io/gorm"
)

type Email struct {
	ID        int64  `gorm:"primaryKey;autoIncrement;not null"`
	To        string `gorm:"not null"`
	Subject   string `gorm:"not null"`
	Body      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (e *Email) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	e.CreatedAt = now
	e.UpdatedAt = now
	return nil
}

func (e *Email) BeforeUpdate(tx *gorm.DB) (err error) {
	e.UpdatedAt = time.Now()
	return nil
}
