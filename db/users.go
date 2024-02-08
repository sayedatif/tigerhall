package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64  `gorm:"primaryKey;autoIncrement;not null"`
	FirstName string `gorm:"type:varchar(255);not null"`
	LastName  string `gorm:"type:varchar(255);not null"`
	Username  string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Tigers    []Tiger `gorm:"many2many:user_tiger_sightings"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return nil
}
