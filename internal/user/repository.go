// package user provides a repository for user
// Add the following functions to the repository
//
// GetById(id uint) (*User, error)
// GetByEmail(email string) (*User, error)
// Create(user User) (*User, error)
// Update(user User) (*User, error)
// Delete(id uint) error

package user

import (
	"context"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetById(c context.Context, id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetByEmail(c context.Context, email string) (*User, error) {
	var user User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) Create(c context.Context, user User) (*User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates a user
func (r *repository) Update(c context.Context, user User) (*User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Delete deletes a user
func (r *repository) Delete(c context.Context, id uint) error {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
