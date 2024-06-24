package controller

import (
	"net/http"
	"strconv"

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
		Message: "result note set",
		Data: map[string]interface{}{
			"rows":       responseArr,
			"total_rows": total,
		}})
}

// GetNote controller
func (n *NoteController) GetNote(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedId, err := strconv.ParseInt(id, 10, 64) //string to integer
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "ID is invalid x-x")
		return
	}
	var note models.Note
	note.ID = parsedId
	foundNote, err := n.service.Find(note)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Finding note went wrong x-x")
		return
	}
	response := foundNote.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of note",
		Data:    &response,
	})
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

// DeleteNote controller
func (n *NoteController) DeleteNote(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "ID is invalid x-x")
		return
	}
	var note models.Note
	note.ID = parsedId
	err = n.service.Delete(parsedId)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "delete note went wrong x-x")
		return
	}

	ctx.JSON(http.StatusOK, util.Response{
		Success: true,
		Message: "Deleted successfully 'U'",
	})
}

// UpdateNote controller
func (n *NoteController) UpdateNote(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "ID is invalid x-x")
		return
	}
	var note models.Note
	note.ID = parsedId

	noteRecord, err := n.service.Find(note)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Note not found x-x")
		return
	}
	ctx.ShouldBindJSON(&noteRecord)

	if noteRecord.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if noteRecord.Description == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Description is required")
	}
	error := n.service.Update(noteRecord)
	if error != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Updating note went wrong x-x")
		return
	}
	response := noteRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Updated successfully 'U'",
		Data:    &response,
	})
}
