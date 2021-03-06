package models

import (
	"gorm.io/gorm"
)

// User struct to store the user data
type User struct {
	gorm.Model
	Username     string `gorm:"unique" json:"username"`
	Password     string `json:"password"`
	FullName     string `json:"fullName"`
	Supervisor   bool   `gorm:"default:0" json:"supervisor"`
	PhoneNumber  string `json:"phoneNumber"`
	ParentNumber string `json:"parentNumber"`
}
