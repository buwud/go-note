package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gonote.com/infrastructure"
)

type note struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

// notes slice
var notes = []note{
	{ID: "1", Title: "odev", Description: "odev detay", Date: time.Now().GoString()},
	{ID: "2", Title: "dev", Description: "kod detay", Date: time.Now().GoString()},
	{ID: "3", Title: "ders notlari", Description: "cok detay", Date: time.Now().GoString()},
}

func GetNotes(c *gin.Context) {
	infrastructure.LoadEnv()
	infrastructure.NewDatabase()
	c.IndentedJSON(http.StatusOK, notes)
}

func GetNoteById(c *gin.Context) {
	id := c.Param("id")
	for _, note := range notes {
		if note.ID == id {
			c.IndentedJSON(http.StatusOK, note)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "note not found"})
}

func DeleteNote(c *gin.Context) {
	id := c.Param("id")
	for i, note := range notes {
		if note.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			c.IndentedJSON(http.StatusOK, notes)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "note not found"})
}

func AddNote(c *gin.Context) {
	var newNote note
	if err := c.BindJSON(&newNote); err != nil {
		return
	}
	notes = append(notes, newNote)
	c.IndentedJSON(http.StatusCreated, newNote)
}

func UpdateNote(c *gin.Context) {
	id := c.Param("id")
	var updatedNote note
	if err := c.BindJSON(&updatedNote); err != nil {
		return
	}
	for _, note := range notes {
		if note.ID == id {
			note = updatedNote
			c.IndentedJSON(http.StatusOK, updatedNote)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "note not found"})
}
