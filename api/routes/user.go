package routes

import (
	"gonote.com/api/controller"
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
		user.POST("/signup", u.Controller.SignUp)
		user.POST("/signin", u.Controller.SignIn)
		user.GET("/users", u.Controller.GetUsers)
	}
	user.GET("/user/notes", u.Controller.GetUserNotes)
}
