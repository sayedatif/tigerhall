package db

import (
	"time"

	"gorm.io/gorm"
)

type UserTigerSighting struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;not null"`
	UserId    uint      `gorm:"not null"`
	TigerId   uint      `gorm:"not null"`
	Lat       float64   `gorm:"not null"`
	Long      float64   `gorm:"not null"`
	SeenAt    time.Time `gorm:"not null"`
	ImageUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Tiger     Tiger `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *UserTigerSighting) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
	return nil
}

func (u *UserTigerSighting) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return nil
}
