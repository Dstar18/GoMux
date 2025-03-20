package repository

import (
	"GoMux/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]entity.User, error)
	GetByID(id uint) (*entity.User, error)
	Create(user *entity.User) error
	GetByField(param string, value interface{}) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id uint) error
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

func (r *userRepository) GetByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByField(param string, value interface{}) (*entity.User, error) {
	var user entity.User
	err := r.db.Where(param+" = ?", value).First(&user).Error
	return &user, err
}

func (r *userRepository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}
