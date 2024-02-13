package tigers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sayedatif/tigerhall/db"
	"github.com/sayedatif/tigerhall/types"
	"github.com/sayedatif/tigerhall/utils"
	"gorm.io/gorm"
)

// @Summary GetTigers
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param page_size query int false "Page size"
// @Success 200 {object} types.TigerResponse
// @Failure 500 {object} types.InternalServerError
// @Router /tigers [get]
func (t TigerController) GetTigers(c *gin.Context) {
	database := t.DB
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

	response := make([]types.TigerResponse, 0)

	for _, t := range tiger {
		response = append(response, types.TigerResponse{
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
