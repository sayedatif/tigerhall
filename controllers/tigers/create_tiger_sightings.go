package tigers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sayedatif/tigerhall/db"
	"github.com/sayedatif/tigerhall/utils"
	"gorm.io/gorm"
)

type CreateTigerSightingBody struct {
	Lat    float64 `json:"lat" binding:"required"`
	Long   float64 `json:"long" binding:"required"`
	SeenAt string  `json:"seen_at" binding:"required"`
}

func (t TigerController) CreateTigerSighting(c *gin.Context) {
	user_id := c.MustGet("user_id")
	var body CreateTigerSightingBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tigerID := c.Param("tiger_id")
	numTigerID, err := strconv.Atoi(tigerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

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

	parsedSeenAt, err := utils.GetParsedTime(body.SeenAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	numUserId := user_id.(float64)

	createTigerSighting := db.UserTigerSighting{UserId: uint(numUserId), TigerId: uint(numTigerID), SeenAt: parsedSeenAt, Lat: body.Lat, Long: body.Long}
	if err := database.Create(&createTigerSighting).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Created tiger sighting successfully",
	})
}
