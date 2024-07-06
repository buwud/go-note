package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gonote.com/api/controller"
	"gonote.com/infrastructure"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type UserRoute struct {
	Controller controller.UserController
	Handler    infrastructure.GinRouter
}

// initialize new route
func NewUserRoute(controller controller.UserController, handler infrastructure.GinRouter) UserRoute {
	return UserRoute{
		Controller: controller,
		Handler:    handler,
	}
}

// setup
func (u UserRoute) Setup() {
	user := u.Handler.Gin.Group("/auth")
	{
		user.POST("/signup", u.Controller.SignUp)
		user.POST("/signin", u.Controller.SignIn)
	}

	// Protected routes
	auth := u.Handler.Gin.Group("/auth")
	auth.Use(authMiddleware())
	{
		auth.GET("/users", u.Controller.GetUsers)
		auth.GET("/user/notes", u.Controller.GetUserNotes)
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", 1)
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

//simdilik burda
