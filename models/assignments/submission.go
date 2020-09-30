package assignments

import (
	usersModel "backend/models/users"
	"time"
)

type AssignmentSubmission struct {
	UploadedAt   time.Time       `json:"uploadedAt"`
	UserID       uint            `json:"userID" gorm:"primary_key;autoIncrement:false"`
	User         usersModel.User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AssignmentID uint            `json:"assignmentID" gorm:"primary_key;autoIncrement:false"`
	Assignment   Assignment      `json:"assignment" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Mark         int             `json:"mark"`
	Graded       bool            `json:"graded" gorm:"default:0"`
	File         string          `json:"file"`
	Feedback     string          `json:"feedback"`
}
