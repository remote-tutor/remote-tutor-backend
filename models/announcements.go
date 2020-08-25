package models

import (
	"github.com/jinzhu/gorm"
)

// Announcement struct to store the announcement data
type Announcement struct {
	gorm.Model
	User    User
	UserID  uint   `json:"user_id"`
	Title   string `json:"title"`
	Topic   string `json:"topic"`
	Content string `json:"content"`
}
