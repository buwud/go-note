package service

import (
	"gonote.com/api/repository"
	"gonote.com/models"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return UserService{
		repository: r,
	}
}

// signup user
func (u UserService) SignUp(user models.UserRegister) error {
	return u.repository.SignUp(user)
}

// signin user
func (u UserService) SignIn(user models.UserLogin) (*models.User, error) {
	return u.repository.SignIn(user)
}
