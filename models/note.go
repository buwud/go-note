package models

import "time"

type Note struct {
	ID          int64     `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `gorm:"size:200" json:"title"`
	Description string    `gorm:"size:500" json:"description"`
	Date        string    `json:"date"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

//sets table name for the model
func (note *Note) TableName() string {
	return "note"
}

//response-map method of model
func (note *Note) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = note.ID
	resp["title"] = note.Title
	resp["description"] = note.Description
	resp["date"] = note.Date
	resp["created_at"] = note.CreatedAt
	resp["updated_at"] = note.UpdatedAt
	return resp
}
