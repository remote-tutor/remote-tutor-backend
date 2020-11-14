package payments

import (
	classesModel "backend/models/organizations"
	"time"

	usersModel "backend/models/users"

	"gorm.io/gorm"
)

// Payment struct to store the users payments data
type Payment struct {
	gorm.Model
	UserID    uint               `json:"userID"`
	User      usersModel.User    `json:"user"`
	StartDate time.Time          `json:"startDate"`
	EndDate   time.Time          `json:"endDate"`
	ClassHash string             `json:"classHash" gorm:"size:25"`
	Class     classesModel.Class `json:"class" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ClassHash;references:Hash"`
}
