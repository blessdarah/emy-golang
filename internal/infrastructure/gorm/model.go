package model

type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	UserID uint   `json:"user_id" gorm:"foreignKey:UserID, references:ID"`
}

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Username     string `json:"username" gorm:"unique"`
	Password     string `json:"-"`
	RefreshToken string `json:"-"`
	Books        []Book `json:"books" gorm:"onetomany"`
}
