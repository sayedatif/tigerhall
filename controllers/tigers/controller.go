package tigers

import "gorm.io/gorm"

type TigerController struct {
	DB *gorm.DB
}
