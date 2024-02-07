package tigers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t TigerController) CreateTigerSighting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Created tiger sighting successfully",
		"data":    "",
	})
}
