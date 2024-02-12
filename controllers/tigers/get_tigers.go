package tigers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sayedatif/tigerhall/db"
	"github.com/sayedatif/tigerhall/utils"
	"gorm.io/gorm"
)

type TigerResponse struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	DOB          string    `json:"dob"`
	LastSeenAt   time.Time `json:"last_seen_at"`
	LastSeenLat  float64   `json:"last_seen_lat"`
	LastSeenLong float64   `json:"last_seen_long"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (t TigerController) GetTigers(c *gin.Context) {
	database := db.GetDB()
	page := c.DefaultQuery("page", "1")
	intPage, _ := utils.StringToInt(page)
	pageSize := c.DefaultQuery("page_size", "10")
	intPageSize, _ := utils.StringToInt(pageSize)

	var tiger []db.Tiger
	if err := database.Order("last_seen_at desc").Limit(intPageSize).Offset((intPage - 1) * intPageSize).Find(&tiger).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	response := make([]TigerResponse, 0)

	for _, t := range tiger {
		response = append(response, TigerResponse{
			ID:           t.ID,
			Name:         t.Name,
			DOB:          t.DOB,
			LastSeenAt:   t.LastSeenAt,
			LastSeenLat:  t.LastSeenLat,
			LastSeenLong: t.LastSeenLong,
			CreatedAt:    t.CreatedAt,
			UpdatedAt:    t.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Fetched tigers successfully",
		"data":    response,
	})
}
