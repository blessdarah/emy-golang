package book

type CreateRequest struct {
	Title  string `json:"title" validate:"min=3,required"`
	Author string `json:"author" validate:"min=3,required"`
	Year   uint   `json:"year" validate:"gt=0"`
}
type UpdateRequest struct {
	Title  *string `json:"title"`
	Author *string `json:"author" validate:"min=3"`
	Year   *uint   `json:"year" validate:"gt=0"`
}

type Response struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   uint   `json:"year"`
}

func (b *CreateRequest) ToModel() Book {
	return Book{
		Title:  b.Title,
		Author: b.Author,
		Year:   b.Year,
	}
}

func (b *Book) ToDto() CreateRequest {
	return CreateRequest{
		Title:  b.Title,
		Author: b.Author,
		Year:   b.Year,
	}
}
