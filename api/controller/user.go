package controller

import (
	"log"
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
const hmacSampleSecret = "AccessToken"

func (u *UserController) SignIn(c *gin.Context) {
	var user models.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid JSON Provided")
		return
	}
	dbUser, err := u.service.SignIn(user)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid Login Credentials")
		return
	}

	// Debugging output
	log.Printf("dbUser: %+v", dbUser)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": dbUser,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		// Debugging output
		log.Printf("Error signing token: %v", err)
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to get token")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Token generated successfully",
		Data:    tokenString,
	}
	c.JSON(http.StatusOK, response)
}

func (u *UserController) GetUserNotes(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		util.ErrorJSON(ctx, http.StatusUnauthorized, "User not found in context")
		return
	}

	// Debugging output to check the type of 'user'
	log.Printf("Type of user in context: %T", user)

	claims, ok := user.(*service.CustomClaim)
	if !ok {
		// Debugging output if the type assertion fails
		log.Printf("Type assertion to *CustomClaim failed, user: %+v", user)
		util.ErrorJSON(ctx, http.StatusUnauthorized, "Invalid token claims")
		return
	}

	username := claims.User.Username // Access the Username from the claims

	// Debugging output
	log.Printf("Username from claims: %s", username)

	notes, err := u.service.GetUserNotes(username)
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
