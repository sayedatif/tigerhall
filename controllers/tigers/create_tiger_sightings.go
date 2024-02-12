package tigers

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sayedatif/tigerhall/db"
	"github.com/sayedatif/tigerhall/utils"
	"gorm.io/gorm"
)

func (t TigerController) CreateTigerSighting(c *gin.Context) {
	user_id := c.MustGet("user_id")
	lat := c.Request.FormValue("lat")
	if lat == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Lat is required"})
		return
	}

	long := c.Request.FormValue("long")
	if long == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Long is required"})
		return
	}

	seenAtStr := c.Request.FormValue("seen_at")
	if seenAtStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "seen_at is required"})
		return
	}

	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}
	defer file.Close()

	ext := filepath.Ext(handler.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported file format."})
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

	numUserId := user_id.(float64)
	filePath, err := utils.HandleImageUpload(file, int(numUserId), tigerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	parsedSeenAt, err := utils.GetParsedTime(seenAtStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	floatLat, _ := strconv.ParseFloat(lat, 64)
	floatLong, _ := strconv.ParseFloat(long, 64)
	createTigerSighting := db.UserTigerSighting{UserId: uint(numUserId), TigerId: uint(numTigerID), SeenAt: parsedSeenAt, Lat: floatLat, Long: floatLong, ImageUrl: filePath}
	if err := database.Create(&createTigerSighting).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Created tiger sighting successfully",
	})
}
