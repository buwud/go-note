package repository

import (
	"gonote.com/infrastructure"
	"gonote.com/models"
)

type NoteRepository struct {
	db infrastructure.Database
}

// new note - fetch db
func NewNoteRepo(db infrastructure.Database) NoteRepository {
	return NoteRepository{
		db: db,
	}
}
func (n NoteRepository) Save(note models.Note) error {
	return n.db.DB.Create(&note).Error
}

func (n NoteRepository) FindAll(note models.Note, keyword string) (*[]models.Note, int64, error) {
	var notes []models.Note
	var totalRows int64 = 0

	queryBuilder := n.db.DB.Order("created_at desc").Model(&models.Note{})

	//search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where(
			n.db.DB.Where("note.title LIKE ? ", queryKeyword))
	}
	err := queryBuilder.
		Where(note).
		Find(&notes).
		Count(&totalRows).Error
	return &notes, totalRows, err
}

//to be continued.....
