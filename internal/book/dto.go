package book

type CreateRequest struct {
	Title  string `json:"title" binding:"required,min=3"`
	Author string `json:"author" binding:"required,min=3"`
	Year   uint   `json:"year" binding:"required,gt=0"`
}
type UpdateRequest struct {
	Title  *string `json:"title" binding:"omitempty,min=3"`
	Author *string `json:"author" binding:"omitempty,min=3"`
	Year   *uint   `json:"year" binding:"omitempty,gt=0"`
}

type Response struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   uint   `json:"year"`
}

func (ur *UpdateRequest) HasUpdates() bool {
	return ur.Title != nil || ur.Author != nil || ur.Year != nil
}

func (ur *UpdateRequest) ApplyToModel(book *Book) {
	if ur.Title != nil {
		book.Title = *ur.Title
	}
	if ur.Author != nil {
		book.Author = *ur.Author
	}
	if ur.Year != nil {
		book.Year = *ur.Year
	}
}

func (b *CreateRequest) ToModel() Book {
	return Book{
		Title:  b.Title,
		Author: b.Author,
		Year:   b.Year,
	}
}
