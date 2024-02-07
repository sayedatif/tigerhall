package tigers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t TigerController) GetTigerSightings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Fetched tiger sightings successfully",
		"data":    "",
	})
}
