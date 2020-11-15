package quizzes

import (
	usersModel "backend/models/users"
	"time"
)

// QuizGrade struct to store the Grade data
type QuizGrade struct {
	Grade     int             `json:"grade"`
	Quiz      Quiz            `json:"quiz" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	QuizID    uint            `json:"quizID" gorm:"primary_key;autoIncrement:false"`
	User      usersModel.User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uint            `json:"userID" gorm:"primary_key;autoIncrement:false"`
	StartAt   time.Time       `json:"startAt"`
	ValidTill time.Time       `json:"validTill"`
}
