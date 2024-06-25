package routes

import (
	"gonote.com/api/controller"
	"gonote.com/infrastructure"
)

type NoteRoute struct {
	Controller controller.NoteController
	Handler    infrastructure.GinRouter
}

// initialize new route
func NewNoteRoute(
	controller controller.NoteController,
	handler infrastructure.GinRouter,
) NoteRoute {
	return NoteRoute{
		Controller: controller,
		Handler:    handler}
}

// setup
func (n NoteRoute) Setup() {
	note := n.Handler.Gin.Group("/notes")
	{
		note.GET("/", n.Controller.GetNotes)
		note.GET("/:id", n.Controller.GetNote)
		note.POST("/", n.Controller.AddNote)
		note.DELETE("/:id", n.Controller.DeleteNote)
		note.PUT("/:id", n.Controller.UpdateNote)
	}
}
