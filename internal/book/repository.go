package book

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewBookRepository(ctx context.Context, db *gorm.DB) *repo {
	return &repo{
		db: db,
	}
}

// handle crud functionality here

func (r *repo) GetAll(c *gin.Context) []Book {
	var books []Book
	r.db.Find(&books)
	return books
}

func (r *repo) GetById(c *gin.Context, id uint) (*Book, error) {
	var book Book
	if err := r.db.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("book %d not found", id)
		}
		return nil, err
	}
	return &book, nil
}
func (r *repo) GetByUserId(c *gin.Context, userId uint) []Book {
	var books []Book
	if err := r.db.Where("user_id = ?", userId).Find(&books).Error; err != nil {
		return nil
	}
	return books
}

func (r *repo) Create(c *gin.Context, book Book) (*Book, error) {
	if err := r.db.Create(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *repo) Update(c *gin.Context, book Book) (*Book, error) {
	if err := r.db.Save(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *repo) Delete(c *gin.Context, id uint) error {
	return r.db.Delete(&Book{}, id).Error
}
