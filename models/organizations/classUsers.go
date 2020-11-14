package organizations

import (
	usersModel "backend/models/users"
	"gorm.io/gorm"
)

// ClassUser holds the users in each class (admin or student)
type ClassUser struct {
	gorm.Model
	UserID    uint            `json:"userID" gorm:"uniqueIndex:idx_user_class,sort:asc"`
	User      usersModel.User `json:"user"`
	ClassHash string          `json:"classHash" gorm:"size:25;uniqueIndex:idx_user_class,sort:asc"`
	Class     Class           `json:"class" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ClassHash;references:Hash"`
	Activated bool            `json:"activated" gorm:"default:0"`
	Admin     bool            `json:"admin" gorm:"default:0"`
}
