package tigers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sayedatif/tigerhall/db"
	"github.com/sayedatif/tigerhall/types"
	"github.com/sayedatif/tigerhall/utils"
)

// @Summary CreateTiger
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer <token>"
// @Param createTiger body types.CreateTigerBody true "Create tiger information"
// @Success 200 {object} types.CreateTiger
// @Failure 500 {object} types.InternalServerError
// @Router /tigers [post]
func (t TigerController) CreateTiger(c *gin.Context) {
	user_id := c.MustGet("user_id")
	var body types.CreateTigerBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parsedLastSeenAt, err := utils.GetParsedTime(body.LastSeenAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	database := t.DB

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

	var response types.CreateTiger
	response.TigerId = createTiger.ID

	c.JSON(http.StatusOK, gin.H{
		"message": "Created new tiger successfully",
		"data":    response,
	})
}
