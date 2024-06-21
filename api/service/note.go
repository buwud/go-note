package service

import (
	"gonote.com/api/repository"
	"gonote.com/models"
)

type NoteService struct {
	repository repository.NoteRepository
}

// return noteService instance
func NewNoteService(r repository.NoteRepository) NoteService {
	return NoteService{
		repository: r,
	}
}

// call --> save repo method
func (n NoteService) Save(note models.Note) error {
	return n.repository.Save(note)
}

// call --> create repo method
func (n NoteService) FindAll(note models.Note, keyword string) (*[]models.Note, int64, error) {
	return n.repository.FindAll(note, keyword)
}

// call --> find repo method
func (n NoteService) Find(note models.Note) (models.Note, error) {
	return n.repository.Find(note)
}

// call --> update repo method
func (n NoteService) Update(note models.Note) error {
	return n.repository.Update(note)
}

// call --> delete repo method
func (n NoteService) Delete(id int64) error {
	var note models.Note
	note.ID = id
	return n.repository.Delete(note)
}
