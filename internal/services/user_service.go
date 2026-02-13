package services

import (
	"context"
	"go_book_api/internal"
	model "go_book_api/internal/infrastructure/gorm"
	"go_book_api/internal/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetAll(c *gin.Context) []model.User
	GetById(c *gin.Context) (*model.User, error)
	Create(c *gin.Context, userReqest internal.UserRequest) (*model.User, error)
	Update(c *gin.Context, userRequest internal.UserRequest) (*model.User, error)
	Delete(c *gin.Context) error
	GetUserBooks(c *gin.Context, userId uint) []model.Book
}

type userService struct {
	ctx      context.Context
	repo     repository.UserRepository
	bookRepo repository.BookRepository
}

func NewUserService(
	ctx context.Context,
	repo repository.UserRepository,
	bookRepo repository.BookRepository,
) *userService {
	return &userService{
		ctx:      ctx,
		repo:     repo,
		bookRepo: bookRepo,
	}
}

func (s *userService) GetAll(c *gin.Context) []model.User {
	return s.repo.GetAll(c)
}

func (s *userService) GetById(c *gin.Context) (*model.User, error) {
	return s.repo.Get(c)
}

func (s *userService) Create(c *gin.Context, userRequest internal.UserRequest) (*model.User, error) {
	// transform user request to user model
	userModel := userRequest.ToModel()

	return s.repo.Create(c, userModel)
}

func (s *userService) Update(c *gin.Context, userRequest internal.UserRequest) (*model.User, error) {
	// transform user request to user model
	userModel := userRequest.ToModel()
	return s.repo.Update(c, userModel)
}

func (s *userService) Delete(c *gin.Context) error {
	return s.repo.Delete(c)
}

func (s *userService) GetUserBooks(c *gin.Context, userId uint) []model.Book {
	return s.bookRepo.GetByUserId(c, userId)
}
