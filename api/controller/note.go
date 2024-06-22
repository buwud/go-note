package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gonote.com/api/service"
	"gonote.com/models"
	"gonote.com/util"
)

type NoteController struct {
	service service.NoteService
}

func NewNoteController(s service.NoteService) NoteController {
	return NoteController{
		service: s,
	}
}

// GetNotes controller
func (n *NoteController) GetNotes(ctx *gin.Context) {
	var notes models.Note
	keyword := ctx.Query("keyword")

	data, total, err := n.service.FindAll(notes, keyword)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find")
		return
	}

	responseArr := make([]map[string]interface{}, 0, 0)

	for _, ni := range *data {
		resp := ni.ResponseMap()
		responseArr = append(responseArr, resp)
	}

	ctx.JSON(http.StatusOK, &util.Response{Success: true,
		Message: "Note post set",
		Data: map[string]interface{}{
			"rows":       responseArr,
			"total_rows": total,
		}})
}

// AddNote controller
func (n *NoteController) AddNote(ctx *gin.Context) {
	var note models.Note
	ctx.ShouldBindJSON(&note)

	if note.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if note.Description == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Description is required")
		return
	}
	err := n.service.Save(note)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to Create Note")
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Note Successfully Created")
}

// GetNote controller will be added
