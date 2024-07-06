package repository

import (
	"gonote.com/infrastructure"
	"gonote.com/models"
	"gonote.com/util"
)

type UserRepository struct {
	db infrastructure.Database
}

func NewUserRepo(db infrastructure.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

// signup user
func (u UserRepository) SignUp(user models.UserRegister) error {
	var newUser models.User
	newUser.FirstName = user.FirstName
	newUser.LastName = user.LastName
	newUser.Username = user.Username
	newUser.Password = user.Password
	newUser.IsActive = true
	return u.db.DB.Create(&newUser).Error
}

// signin user
func (u UserRepository) SignIn(user models.UserLogin) (*models.User, error) {
	var dbUser models.User
	username := user.Username
	pass := user.Password
	err := u.db.DB.Where("username = ?", username).First(&dbUser).Error
	if err != nil {
		return nil, err
	}

	passErr := util.CheckPasswordHash(pass, dbUser.Password)
	if passErr != nil {
		return nil, passErr
	}
	return &dbUser, nil
}

// get users
func (u UserRepository) GetUsers() (*[]models.User, error) {
	var users []models.User
	err := u.db.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

// get user's notes
func (u UserRepository) GetUserNotes(userid int64) (*[]models.Note, error) {
	var notes []models.Note
	err := u.db.
		DB.
		Where("user_id = ?", userid).
		Find(&notes).
		Error
	if err != nil {
		return nil, err
	}
	return &notes, nil
}
