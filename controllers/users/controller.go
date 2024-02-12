package users

import "gorm.io/gorm"

type UserController struct {
	DB *gorm.DB
}
