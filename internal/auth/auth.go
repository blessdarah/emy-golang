package auth

import (
	"net/http"

	"go_book_api/config"
	"go_book_api/internal/model"
	"go_book_api/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		model.ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	var existing model.User
	if err := config.DB.Where("username = ?", user.Username).First(&existing).Error; err == nil {
		model.ResponseJSON(c, http.StatusConflict, "User already exists", nil)
		return
	}

	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	config.DB.Create(&user)
	model.ResponseJSON(c, http.StatusCreated, "User created successfully", user)
}

func Login(c *gin.Context) {
	var input model.User
	var user model.User

	if err := c.ShouldBindJSON(&input); err != nil {
		model.ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		model.ResponseJSON(c, http.StatusUnauthorized, "No such user", nil)
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		model.ResponseJSON(c, http.StatusUnauthorized, "Invalid password", nil)
		return
	}

	token, _ := utils.GenerateJWT(user.Username)
	refreshToken, _ := utils.GenerateRefreshToken(user.Username)

	user.RefreshToken = refreshToken
	config.DB.Save(&user)

	model.ResponseJSON(c, http.StatusOK, "Login successful", []string{token})
}

func Logout(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		model.ResponseJSON(c, http.StatusBadRequest, "No username in context", nil)
		return
	}
	uname, ok := username.(string)
	if !ok {
		model.ResponseJSON(c, http.StatusBadRequest, "Invalid username", nil)
		return
	}

	if err := config.DB.Model(&model.User{}).
		Where("username = ?", uname).
		Update("refresh_token", "").Error; err != nil {
		model.ResponseJSON(c, http.StatusInternalServerError, "Failed to logout", nil)
		return
	}

	model.ResponseJSON(c, http.StatusOK, "Logged out successfully", nil)
}

func RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		model.ResponseJSON(c, http.StatusBadRequest, "Bad Request", nil)
		return
	}

	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return utils.RefreshSecret, nil
	})
	if err != nil || !token.Valid {
		model.ResponseJSON(c, http.StatusUnauthorized, "Invalid refresh token", nil)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	var user model.User
	config.DB.Where("username = ?", username).First(&user)
	if user.RefreshToken != req.RefreshToken {
		model.ResponseJSON(c, http.StatusUnauthorized, "Token mismatch", nil)
		return
	}

	newToken, _ := utils.GenerateJWT(username)
	model.ResponseJSON(c, http.StatusOK, "New access token", newToken)
}
