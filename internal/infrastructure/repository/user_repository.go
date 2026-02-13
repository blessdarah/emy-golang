package repository

import (
	"context"
	model "go_book_api/internal/infrastructure/gorm"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll(c *gin.Context) []model.User
	Get(c *gin.Context) (*model.User, error)
	Create(c *gin.Context, user model.User) (*model.User, error)
	Update(c *gin.Context, user model.User) (*model.User, error)
	Delete(c *gin.Context) error
}

type userRepository struct {
	ctx context.Context
	db  *gorm.DB
}

func NewUserRepository(ctx context.Context, db *gorm.DB) *userRepository {
	return &userRepository{
		ctx: ctx,
		db:  db,
	}
}

func (r *userRepository) GetAll(c *gin.Context) []model.User {
	var users []model.User
	r.db.Find(&users)
	return users
}

func (r *userRepository) Get(c *gin.Context) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, c.Param("id")).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(c *gin.Context, user model.User) (*model.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(c *gin.Context, user model.User) (*model.User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Delete(c *gin.Context) error {
	return r.db.Delete(&model.User{}, c.Param("id")).Error
}
