package usecase

import (
	"GoMux/entity"
	"GoMux/repository"
)

type UserUseCase interface {
	GetUsers() ([]entity.User, error)
	GetUserById(id uint) (*entity.User, error)
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
