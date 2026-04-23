package book

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID     uint `gorm:"primaryKey"`
	Title  string
	Author string
	Year   uint
	UserID uint `gorm:"foreignKey:UserID, references:ID"`
}

// BeforeCreate hooks to generate id bofore
// inserting into database
func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uint(time.Now().Unix())
	return
}

// ToResponse transforms book model to response
func (b *Book) ToResponse() Response {
	return Response{
		ID:     b.ID,
		Title:  b.Title,
		Author: b.Author,
		Year:   b.Year,
	}
}
