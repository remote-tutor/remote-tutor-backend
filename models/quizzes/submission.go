package quizzes

import (
	md "backend/models"
)

// Submission struct to store the Submission data
type Submission struct {
	Grade      int      `json:"grade"`
	Question   Question `json:"question"`
	QuestionID uint     `json:"questionID" gorm:"primary_key;auto_increment:false"`
	User       md.User  `json:"user"`
	UserID     uint     `json:"userID" gorm:"primary_key;auto_increment:false"`
}

// MCQSubmission struct to store the MCQ Submission type data
type MCQSubmission struct {
	Submission `json:"submission"`
	UserResult uint `json:"userResult"`
}

// LongAnswerSubmission struct to store the LongAnswer Submission type data
type LongAnswerSubmission struct {
	Submission `json:"submission"`
	UserResult string `json:"userResult"`
}
