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
	noteRepository := repository.NewNoteRepo(db)
	noteService := service.NewNoteService(noteRepository)
	noteController := controller.NewNoteController(noteService)
	noteRoute := routes.NewNoteRoute(noteController, router)
	noteRoute.Setup()

	db.DB.AutoMigrate(&models.Note{})
	router.Gin.Run(":8000")
}
