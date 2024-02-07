package tigers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t TigerController) CreateTiger(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Created new tiger successfully",
		"data":    "",
	})
}
