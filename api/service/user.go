package service

import (
	"github.com/golang-jwt/jwt"
	"gonote.com/api/repository"
	"gonote.com/models"
)

type CustomClaim struct {
	User models.User `json:"user"`
	jwt.StandardClaims
}

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

// get user's notes
func (u UserService) GetUserNotes(userid int64) (*[]models.Note, error) {
	return u.repository.GetUserNotes(userid)
}

// get users
func (u UserService) GetUsers() (*[]models.User, error) {
	return u.repository.GetUsers()
}
