package tigers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sayedatif/tigerhall/db"
)

type CreateTigerBody struct {
	Name         string  `json:"name" binding:"required"`
	DOB          string  `json:"dob" binding:"required"`
	LastSeenAt   string  `json:"last_seen_at" binding:"required"`
	LastSeenLat  float64 `json:"last_seen_lat" binding:"required"`
	LastSeenLong float64 `json:"last_seen_long" binding:"required"`
}

func (t TigerController) CreateTiger(c *gin.Context) {
	var body CreateTigerBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := db.GetDB()

	layout := "2006-01-02 15:04:05.999"

	parsedLastSeenAt, err := time.Parse(layout, body.LastSeenAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	createTiger := db.Tiger{Name: body.Name, DOB: body.DOB, LastSeenAt: parsedLastSeenAt, LastSeenLat: body.LastSeenLat, LastSeenLong: body.LastSeenLong}
	now := time.Now()
	createTiger.CreatedAt = now
	createTiger.UpdatedAt = now
	if err := database.Create(&createTiger).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Created new tiger successfully",
	})
}
