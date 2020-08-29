package models

import (
	usersModel "backend/models/users"

	"github.com/jinzhu/gorm"
)

// Announcement struct to store the announcement data
type Announcement struct {
	gorm.Model
	User    usersModel.User `json:"user"`
	UserID  uint            `json:"user_id"`
	Title   string          `json:"title"`
	Topic   string          `json:"topic"`
	Content string          `json:"content"`
}
