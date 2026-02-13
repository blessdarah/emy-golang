package helpers

import (
	"go_book_api/config"
	model "go_book_api/internal/infrastructure/gorm"

	"github.com/gin-gonic/gin"
)

func GetAuthUser(c *gin.Context) (*model.User, error) {
	var user model.User
	username := c.Value("username").(string)
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
