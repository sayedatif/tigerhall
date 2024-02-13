package users

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sayedatif/tigerhall/db"
	"github.com/sayedatif/tigerhall/types"
	"github.com/sayedatif/tigerhall/utils"
	"gorm.io/gorm"
)

// @Summary Login
// @Accept json
// @Produce json
// @Param login body types.LoginBody true "User login information"
// @Success 200 {object} types.LoginResponse
// @Failure 500 {object} types.InternalServerError
// @Router /users/login [post]
func (u UserController) Login(c *gin.Context) {
	var body types.LoginBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database := u.DB
	var user db.User
	if err := database.Where("email = ?", body.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	match := utils.CheckPasswordHash(body.Password, user.Password)
	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	secret := utils.GetEnv("SECRET_KEY")
	byteSecret, err := json.Marshal(secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	token, err := utils.GenerateToken(user.ID, byteSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
		"data": types.LoginResponse{
			Token: token,
		},
	})
}
