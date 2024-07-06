package routes

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gonote.com/api/controller"
	"gonote.com/api/service"
	"gonote.com/infrastructure"
)

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
		user.POST("/register", u.Controller.SignUp)
		user.POST("/login", u.Controller.SignIn)
	}
	// auth middleware
	user.Use(authMiddleware())
	{
		user.GET("/profile", u.Controller.GetUserNotes)
	}
}

const hmacSampleSecret = "AccessToken"

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Authorization Header Not Found"})
			return
		}

		splitToken := strings.Split(auth, "Bearer ")
		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Invalid Authorization Header"})
			return
		}

		auth = splitToken[1]

		token, err := jwt.ParseWithClaims(auth, &service.CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(hmacSampleSecret), nil
		})

		if err != nil {
			var verr *jwt.ValidationError
			if errors.As(err, &verr) {
				switch {
				case verr.Errors&jwt.ValidationErrorMalformed != 0:
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Token is malformed"})
				case verr.Errors&jwt.ValidationErrorExpired != 0:
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Token is expired"})
				case verr.Errors&jwt.ValidationErrorNotValidYet != 0:
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Token not active yet"})
				case verr.Errors&jwt.ValidationErrorSignatureInvalid != 0:
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Invalid token signature"})
				default:
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Token parsing error", "Error": err.Error()})
				}
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Token parsing error", "Error": err.Error()})
			}
			return
		}

		if claims, ok := token.Claims.(*service.CustomClaim); ok && token.Valid {
			c.Set("user", claims)
			//log claims
			log.Printf("CLAIMS: ", claims)
			log.Printf("User: %v, ExpiresAt: %v", claims.User.Username, claims.StandardClaims.ExpiresAt)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Invalid Token Claims"})
			return
		}

		c.Next()
	}
}
