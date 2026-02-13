package repository

import (
	"context"
	model "go_book_api/internal/infrastructure/gorm"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetAll(c *gin.Context) []model.Book
	GetById(c *gin.Context, id uint) (*model.Book, error)
	GetByUserId(c *gin.Context, userId uint) []model.Book
	Create(c *gin.Context, book model.Book) (*model.Book, error)
	Update(c *gin.Context, book model.Book) (*model.Book, error)
	Delete(c *gin.Context, id uint) error
}

type bookRepository struct {
	ctx context.Context
	db  *gorm.DB
}

func NewBookRepository(ctx context.Context, db *gorm.DB) *bookRepository {
	return &bookRepository{
		ctx: ctx,
		db:  db,
	}
}

func (r *bookRepository) GetAll(c *gin.Context) []model.Book {
	var books []model.Book
	r.db.Find(&books)
	return books
}

func (r *bookRepository) GetById(c *gin.Context, id uint) (*model.Book, error) {
	var book model.Book
	if err := r.db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}
func (r *bookRepository) GetByUserId(c *gin.Context, userId uint) []model.Book {
	var books []model.Book
	if err := r.db.Where("user_id = ?", userId).Find(&books).Error; err != nil {
		return nil
	}
	return books
}

func (r *bookRepository) Create(c *gin.Context, book model.Book) (*model.Book, error) {
	if err := r.db.Create(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) Update(c *gin.Context, book model.Book) (*model.Book, error) {
	if err := r.db.Save(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) Delete(c *gin.Context, id uint) error {
	return r.db.Delete(&model.Book{}, id).Error
}
