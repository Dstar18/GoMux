package repository

import (
	"GoMux/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// initialize userrepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAll() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error
	return users, err
}
