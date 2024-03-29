package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sayedatif/tigerhall/db"
	"github.com/sayedatif/tigerhall/types"
	"github.com/sayedatif/tigerhall/utils"
	"gorm.io/gorm"
)

// @Summary Signup
// @Accept json
// @Produce json
// @Param signup body types.SignupBody true "User signup information"
// @Success 200 {object} types.SignupResponse
// @Failure 500 {object} types.InternalServerError
// @Router /users/signup [post]
func (u UserController) Signup(c *gin.Context) {
	var body types.SignupBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := u.DB
	var user db.User
	if err := database.Where("email = ?", body.Email).First(&user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
	}

	if user.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	hashedPassword, _ := utils.HashPassword(body.Password)
	username := utils.GenerateUsername(body.FirstName, body.LastName)

	createUser := db.User{FirstName: body.FirstName, LastName: body.LastName, Email: body.Email, Password: hashedPassword, Username: username}
	if err := database.Create(&createUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User signed up successfully",
	})
}
