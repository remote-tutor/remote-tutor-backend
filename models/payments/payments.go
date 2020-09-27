package payments

import (
	"time"

	usersModel "backend/models/users"

	"gorm.io/gorm"
)

// Payment struct to store the users payments data
type Payment struct {
	gorm.Model
	UserID    uint            `json:"userID"`
	User      usersModel.User `json:"user"`
	StartDate time.Time       `json:"startDate"`
	EndDate   time.Time       `json:"endDate"`
}
