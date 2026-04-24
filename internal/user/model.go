package user

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"username" gorm:"unique,max:75"`
	Password string `json:"password"`
	Email    string `json:"first_name" gorm:"unique"`
}
