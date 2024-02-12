package tigers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sayedatif/tigerhall/db"
	"github.com/sayedatif/tigerhall/utils"
	"gorm.io/gorm"
)

type Result struct {
	Email        string  `json:"email"`
	TigerName    string  `json:"tiger_name"`
	LastSeenLat  float64 `json:"last_seen_lat"`
	LastSeenLong float64 `json:"last_seen_long"`
}

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

	database := t.DB

	tx := database.Begin()
	var tiger db.Tiger
	if err := tx.Where("id = ?", tigerID).First(&tiger).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	floatLat, _ := strconv.ParseFloat(lat, 64)
	floatLong, _ := strconv.ParseFloat(long, 64)
	distance := utils.CalculateDistance(tiger.LastSeenLat, tiger.LastSeenLong, floatLat, floatLong)
	if distance < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tiger is previously sighted within 5 kilometres"})
		return
	}

	numUserId := user_id.(float64)
	filePath, err := utils.HandleImageUpload(file, int(numUserId), tigerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	seenAt := time.Now()
	createTigerSighting := db.UserTigerSighting{UserId: uint(numUserId), TigerId: uint(numTigerID), SeenAt: seenAt, Lat: floatLat, Long: floatLong, ImageUrl: filePath}
	if err := tx.Create(&createTigerSighting).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if err := tx.Model(&db.Tiger{}).Where("id = ?", numTigerID).Updates(db.Tiger{LastSeenAt: seenAt, LastSeenLat: floatLat, LastSeenLong: floatLong}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	var result []Result
	if err := tx.Raw("SELECT u.email, t.name as tiger_name, t.last_seen_lat, t.last_seen_long FROM user_tiger_sightings uts JOIN users u on u.id = uts.user_id JOIN tigers t on t.id = uts.tiger_id WHERE tiger_id = ? and user_id != ? group by uts.user_id, u.email, t.name, t.last_seen_lat, t.last_seen_long", tigerID, numUserId).Scan(&result).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	emailQueue := make(chan utils.Email, 100)
	go func() {
		for {
			email := <-emailQueue
			utils.SendEmail(database, email)
		}
	}()

	for i := 0; i < len(result); i++ {
		email := utils.Email{
			To:      result[i].Email,
			Subject: "Tiger Sighted",
			Body:    fmt.Sprintf("%s has been sighted at %.2f, %.2f", result[i].TigerName, result[i].LastSeenLat, result[i].LastSeenLong),
		}
		emailQueue <- email
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Created tiger sighting successfully",
	})
}
