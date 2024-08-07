package main

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
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

	// Determine the base path of the current file
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	// Serve static files
	router.Gin.Static("/static", filepath.Join(basePath, "build", "static"))

	// Serve index.html for all other routes
	router.Gin.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(basePath, "build", "index.html"))
	})

	//log paths
	log.Printf("Base path: %s", basePath)

	router.Gin.Run(":8000")
}
