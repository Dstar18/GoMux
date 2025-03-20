package usecase

import (
	"GoMux/entity"
	"GoMux/repository"
)

type UserUseCase interface {
	GetUsers() ([]entity.User, error)
	GetUserById(id uint) (*entity.User, error)
	CreateUser(user *entity.User) error
	GetUserByField(param string, value interface{}) (*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id uint) error
}

type userUseCase struct {
	userRepo repository.UserRepository
}

// initialize userUseCase
func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{userRepo}
}

func (u *userUseCase) GetUsers() ([]entity.User, error) {
	return u.userRepo.GetAll()
}

func (u *userUseCase) GetUserById(id uint) (*entity.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *userUseCase) CreateUser(user *entity.User) error {
	return u.userRepo.Create(user)
}

func (u *userUseCase) GetUserByField(param string, value interface{}) (*entity.User, error) {
	return u.userRepo.GetByField(param, value)
}

func (u *userUseCase) UpdateUser(user *entity.User) error {
	return u.userRepo.Update(user)
}

func (u *userUseCase) DeleteUser(id uint) error {
	return u.userRepo.Delete(id)
}
