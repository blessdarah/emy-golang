package handler

import (
	"net/http"
	"strconv"

	"go_book_api/internal/services"
	"go_book_api/utils"

	"github.com/gin-gonic/gin"
)

type UserHanlder interface {
	GetUser(c *gin.Context)
	GetUserBooks(c *gin.Context)
}

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) GetUser(c *gin.Context) {
	userCollection := h.userService.GetAll(c)
	utils.ResponseJSON(
		c,
		http.StatusOK,
		"User retrieved successfully",
		userCollection,
	)
}

func (h *userHandler) GetUserBooks(c *gin.Context) {
	// get user id from url
	id := c.Param("id")
	uintId, _ := strconv.Atoi(id)
	userBookCollection := h.userService.GetUserBooks(c, uint(uintId))
	utils.ResponseJSON(c, http.StatusOK,
		"User books retrieved successfully", userBookCollection)
}
