package quizzes

import (
	usersModel "backend/models/users"
)

// Submission struct to store the Submission data
type Submission struct {
	Grade  int             `json:"grade"`
	User   usersModel.User `json:"user"`
	UserID uint            `json:"userID" gorm:"primary_key;autoIncrement:false"`
}

// MCQSubmission struct to store the MCQ Submission type data
type MCQSubmission struct {
	MCQ        MCQ  `json:"mcq"`
	MCQID      uint `json:"mcqID" gorm:"primary_key;autoIncrement:false"`
	Submission `json:"submission"`
	UserResult uint `json:"userResult"`
}

// LongAnswerSubmission struct to store the LongAnswer Submission type data
type LongAnswerSubmission struct {
	LongAnswer   LongAnswer `json:"longAnswer"`
	LongAnswerID uint       `json:"longAnswerID" gorm:"primary_key;autoIncrement:false"`
	Submission   `json:"submission"`
	UserResult   string `json:"userResult"`
}
