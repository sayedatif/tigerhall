package tigers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sayedatif/tigerhall/db"
	"github.com/sayedatif/tigerhall/utils"
	"gorm.io/gorm"
)

type TigerSightingsResponse struct {
	ID        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	TigerId   int64     `json:"tiger_id"`
	SeenAt    time.Time `json:"seen_at"`
	Lat       float64   `json:"lat"`
	Long      float64   `json:"long"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetTigerSightings(c *gin.Context) {
	tigerID := c.Param("tiger_id")
	database := db.GetDB()
	var tiger db.Tiger
	if err := database.Where("id = ?", tigerID).First(&tiger).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	page := c.DefaultQuery("page", "1")
	intPage, _ := utils.StringToInt(page)
	pageSize := c.DefaultQuery("page_size", "10")
	intPageSize, _ := utils.StringToInt(pageSize)

	var userTigerSighting []db.UserTigerSighting
	if err := database.Where("tiger_id = ?", tigerID).Order("seen_at desc").Limit(intPageSize).Offset((intPage - 1) * intPageSize).Find(&userTigerSighting).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	response := make([]TigerSightingsResponse, 0)

	for _, t := range userTigerSighting {
		response = append(response, TigerSightingsResponse{
			ID:        t.ID,
			UserId:    int64(t.UserId),
			TigerId:   int64(t.TigerId),
			Lat:       t.Lat,
			Long:      t.Long,
			SeenAt:    t.SeenAt,
			ImageUrl:  t.ImageUrl,
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Fetched tiger sightings successfully",
		"data":    response,
	})
}
