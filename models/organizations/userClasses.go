package organizations

import (
	usersModel "backend/models/users"
	"gorm.io/gorm"
)

type UserClass struct {
	gorm.Model
	UserID    uint   `json:"userID" gorm:"uniqueIndex:idx_user_class,sort:asc"`
	User      usersModel.User
	ClassHash string `json:"classHash" gorm:"size:255;uniqueIndex:idx_user_class,sort:asc"`
	Class     Class `json:"class" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ClassHash;references:Hash"`
}
