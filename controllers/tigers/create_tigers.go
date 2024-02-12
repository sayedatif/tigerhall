package tigers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sayedatif/tigerhall/db"
	"github.com/sayedatif/tigerhall/utils"
)

type CreateTigerBody struct {
	Name         string  `json:"name" binding:"required"`
	DOB          string  `json:"dob" binding:"required"`
	LastSeenAt   string  `json:"last_seen_at" binding:"required"`
	LastSeenLat  float64 `json:"last_seen_lat" binding:"required"`
	LastSeenLong float64 `json:"last_seen_long" binding:"required"`
}

func (t TigerController) CreateTiger(c *gin.Context) {
	user_id := c.MustGet("user_id")
	var body CreateTigerBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parsedLastSeenAt, err := utils.GetParsedTime(body.LastSeenAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	database := db.GetDB()

	tx := database.Begin()

	createTiger := db.Tiger{Name: body.Name, DOB: body.DOB, LastSeenAt: parsedLastSeenAt, LastSeenLat: body.LastSeenLat, LastSeenLong: body.LastSeenLong}
	if err := tx.Create(&createTiger).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	numUserId := user_id.(float64)
	createTigerSighting := db.UserTigerSighting{UserId: uint(numUserId), TigerId: uint(createTiger.ID), SeenAt: parsedLastSeenAt, Lat: body.LastSeenLat, Long: body.LastSeenLong}
	if err := tx.Create(&createTigerSighting).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Created new tiger successfully",
	})
}
