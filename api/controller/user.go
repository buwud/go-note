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
	SampleSecret := []byte("your_secret_key") // Use a consistent secret key

	if err := ctx.ShouldBindJSON(&user); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Invalid JSON")
		return
	}

	dbUser, err := u.service.SignIn(user)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Invalid Login")
		return
	}

	// Generate token with user_id claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": dbUser.ID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString(SampleSecret)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to generate token")
		return
	}

	response := &util.Response{
		Success: true,
		Message: "Token generation is successful",
		Data:    tokenString,
	}
	ctx.JSON(http.StatusOK, response)
}
func (u *UserController) GetUserNotes(ctx *gin.Context) {
	claims, exists := ctx.Get("claims")
	if !exists {
		util.ErrorJSON(ctx, http.StatusUnauthorized, "Unauthorized: No claims found in context")
		return
	}

	claimsData, ok := claims.(*routes.Claims)
	if !ok {
		util.ErrorJSON(ctx, http.StatusUnauthorized, "Unauthorized: Invalid token claims")
		return
	}

	notes, err := u.service.GetUserNotes(claimsData.Username)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to get notes")
		return
	}

	response := make([]map[string]interface{}, 0, len(*notes))
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
