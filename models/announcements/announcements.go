package models

import (
	classesModel "backend/models/organizations"
	usersModel "backend/models/users"

	"gorm.io/gorm"
)

// Announcement struct to store the announcement data
type Announcement struct {
	gorm.Model
	User      usersModel.User    `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uint               `json:"userID"`
	Title     string             `json:"title"`
	Topic     string             `json:"topic"`
	Content   string             `json:"content"`
	Year      int                `json:"year"`
	ClassHash string             `json:"classHash" gorm:"size:25"`
	Class     classesModel.Class `json:"class" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ClassHash;references:Hash"`
}
