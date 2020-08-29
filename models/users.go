package models

import (
	"backend/utils"

	"github.com/jinzhu/gorm"
)

// User struct to store the user data
type User struct {
	gorm.Model
	Username  string `gorm:"unique" json:"username"`
	Password  string `json:"password"`
	FullName  string `json:"full_name"`
	Activated bool   `gorm:"default:0" json:"activated"`
	Admin     bool   `gorm:"default:0" json:"admin"`
	Hash      string `json:"hash"`
	Year      int    `json:"year"`
}

// AfterCreate updates the Hash column of the user after creation
func (user *User) AfterCreate(scope *gorm.Scope) error {
	ID := int(user.ID)
	generatedHash := utils.GenerateHash(ID)
	scope.DB().Model(user).Updates(User{Hash: generatedHash})
	return nil
}
