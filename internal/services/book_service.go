package services

import (
	"context"
	"go_book_api/internal"
	model "go_book_api/internal/infrastructure/gorm"
	"go_book_api/internal/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

type BookService interface {
	GetAll(c *gin.Context) []model.Book
	GetById(c *gin.Context, id uint) (*model.Book, error)
	Create(c *gin.Context, bookRequest internal.BookRequest) (*model.Book, error)
	Update(c *gin.Context, bookRequest internal.BookRequest) (*model.Book, error)
	Delete(c *gin.Context, id uint) error
}

type bookService struct {
	ctx  context.Context
	repo repository.BookRepository
}

func NewBookService(ctx context.Context, repo repository.BookRepository) *bookService {
	return &bookService{
		ctx:  ctx,
		repo: repo,
	}
}

func (s *bookService) GetAll(c *gin.Context) []model.Book {
	return s.repo.GetAll(c)
}

func (s *bookService) GetById(c *gin.Context, id uint) (*model.Book, error) {
	return s.repo.GetById(c, id)
}

func (s *bookService) Create(c *gin.Context, bookRequest internal.BookRequest) (*model.Book, error) {
	// transform book request to book model
	bookModel := bookRequest.ToModel()

	return s.repo.Create(c, bookModel)
}

func (s *bookService) Update(c *gin.Context, bookRequest internal.BookRequest) (*model.Book, error) {
	// transform book request to book model
	bookModel := bookRequest.ToModel()
	return s.repo.Update(c, bookModel)
}

func (s *bookService) Delete(c *gin.Context, id uint) error {
	return s.repo.Delete(c, id)
}
