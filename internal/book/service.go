package book

import (
	"context"

	"github.com/gin-gonic/gin"
)

type bookRepository interface {
	GetAll(c *gin.Context) []Book
	GetById(c *gin.Context, id uint) (*Book, error)
	Create(c *gin.Context, book Book) (*Book, error)
	Update(c *gin.Context, book Book) (*Book, error)
	Delete(c *gin.Context, id uint) error
}

type bookService struct {
	ctx  context.Context
	repo bookRepository
}

func NewBookService(ctx context.Context, repo bookRepository) *bookService {
	return &bookService{
		ctx:  ctx,
		repo: repo,
	}
}

func (s *bookService) GetAll(c *gin.Context) []Book {
	return s.repo.GetAll(c)
}

func (s *bookService) GetById(c *gin.Context, id uint) (*Book, error) {
	return s.repo.GetById(c, id)
}

func (s *bookService) Create(c *gin.Context, b Book) (*Book, error) {
	// transform book request to book model
	return s.repo.Create(c, b)
}

func (s *bookService) Update(c *gin.Context, b Book) (*Book, error) {
	// transform book request to book model
	return s.repo.Update(c, b)
}

func (s *bookService) Delete(c *gin.Context, id uint) error {
	return s.repo.Delete(c, id)
}
