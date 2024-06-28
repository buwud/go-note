package models

import "time"

type User struct {
	ID        int64     `gorm:"primary_key:auto_increment" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	IsActive  string    `json:"is_active"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (user *User) TableName() string {
	return "user"
}

//user login struct
type UserLogin struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

//user register struct
type UserRegister struct {
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
}

//response map
func (user *User) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = user.ID
	resp["first_name"] = user.FirstName
	resp["last_name"] = user.LastName
	resp["username"] = user.Username
	resp["is_active"] = user.IsActive
	resp["created_at"] = user.CreatedAt
	resp["updated_at"] = user.UpdatedAt
	return resp
}
