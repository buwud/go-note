package main

import (
	"gonote.com/api/controller"
	"gonote.com/api/repository"
	"gonote.com/api/routes"
	"gonote.com/api/service"
	"gonote.com/infrastructure"
	"gonote.com/models"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {
	// initialize everything
	router := infrastructure.NewGinRouter()
	db := infrastructure.NewDatabase()
	//note
	noteRepository := repository.NewNoteRepo(db)
	noteService := service.NewNoteService(noteRepository)
	noteController := controller.NewNoteController(noteService)
	noteRoute := routes.NewNoteRoute(noteController, router)
	noteRoute.Setup()

	//user
	userRepository := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()

	db.DB.AutoMigrate(&models.Note{}, &models.User{})
	router.Gin.Run(":8000")
}
