package quizzes

import (
	usersModel "backend/models/users"
)

// QuizGrade struct to store the Grade data
type QuizGrade struct {
	Grade  int             `json:"grade"`
	Quiz   Quiz            `json:"quiz"`
	QuizID uint            `json:"quizID" gorm:"primary_key;auto_increment:false"`
	User   usersModel.User `json:"user"`
	UserID uint            `json:"userID" gorm:"primary_key;auto_increment:false"`
}
