package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gonote.com/api/service"
	"gonote.com/models"
	"gonote.com/util"
)

type UserController struct {
	service service.UserService
}

func NewUserController(s service.UserService) UserController {
	return UserController{
		service: s,
	}
}

// signup controller
func (u *UserController) SignUp(ctx *gin.Context) {
	var user models.UserRegister
	if err := ctx.ShouldBind(&user); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Invalid JSON")
		return
	}
	hashPassword, _ := util.HashPassword(user.Password)
	user.Password = hashPassword

	err := u.service.SignUp(user)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to sign-up")
		return
	}
	util.SuccessJSON(ctx, http.StatusOK, "Successfully signed up <3")
}

// signin user
// generate jwt token if user logged in to the system
func (u *UserController) SignIn(ctx *gin.Context) {
	var user models.UserLogin
	var SampleSecret []byte

	if err := ctx.ShouldBindJSON(&user); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Invalid Json")
		return
	}
	dbUser, err := u.service.SignIn(user)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Invalid Login")
		return
	}
	//generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": dbUser,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString(SampleSecret)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to get token")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Token generation is successfull",
		Data:    tokenString,
	}
	ctx.JSON(http.StatusOK, response)
}

// get user's notes by specicific user id, use userid from jwt token
func (u *UserController) GetUserNotes(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	claims := user.(jwt.MapClaims)
	userID := claims["user"].(models.User).ID
	notes, err := u.service.GetUserNotes(models.User{ID: userID})
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to get notes")
		return
	}
	response := make([]map[string]interface{}, 0)
	for _, note := range *notes {
		response = append(response, note.ResponseMap())
	}
	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of notes",
		Data:    &response,
	})
}

// GetUsers controller
func (u *UserController) GetUsers(ctx *gin.Context) {
	users, err := u.service.GetUsers()
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to get users")
		return
	}
	response := make([]map[string]interface{}, 0)
	for _, user := range *users {
		response = append(response, user.ResponseMap())
	}
	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of users",
		Data:    &response,
	})
}
