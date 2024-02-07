package tigers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t TigerController) GetTigers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Fetched tigers successfully",
		"data":    "",
	})
}
