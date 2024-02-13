package tigers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sayedatif/tigerhall/db"
	"github.com/sayedatif/tigerhall/types"
	"github.com/sayedatif/tigerhall/utils"
	"gorm.io/gorm"
)

func (t TigerController) GetTigerSightings(c *gin.Context) {
	tigerID := c.Param("tiger_id")
	database := t.DB
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

	response := make([]types.TigerSightingsResponse, 0)

	for _, t := range userTigerSighting {
		response = append(response, types.TigerSightingsResponse{
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
