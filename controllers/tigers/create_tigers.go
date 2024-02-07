package tigers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t TigerController) CreateTiger(c *gin.Context) {
	user_id := c.MustGet("user_id")

	c.JSON(http.StatusOK, gin.H{
		"message": "Created new tiger successfully",
		"data":    user_id,
	})
}
