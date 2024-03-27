// File: repository/user.go
package repository

import (
	"gorm.io/gorm"

	"github.com/catherineelfrida/mygram-api/internal/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(userID string, updatedUser *model.User) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&user).Updates(updatedUser).Error; err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (r *UserRepository) Delete(userID string) error {
	var user model.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}