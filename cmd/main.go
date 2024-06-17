package main

import (
	"gonote.com/api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/notes", handler.GetNotes)
	router.POST("/notes/:id", handler.GetNoteById)
	router.DELETE("/notes/:id", handler.DeleteNote)
	router.Run("localhost:4242")
}
