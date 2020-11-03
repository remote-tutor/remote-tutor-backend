package quizzes

import (
	submissionsHook "backend/hooks/quizzes"
	usersModel "backend/models/users"
	"time"

	"gorm.io/gorm"
)

// Submission struct to store the Submission data
type Submission struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Grade     int             `json:"grade"`
	User      usersModel.User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uint            `json:"userID" gorm:"primary_key;autoIncrement:false"`
}

// MCQSubmission struct to store the MCQ Submission type data
type MCQSubmission struct {
	MCQ        MCQ  `json:"mcq" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MCQID      uint `json:"mcqID" gorm:"primary_key;autoIncrement:false"`
	Submission `json:"submission"`
	UserResult uint   `json:"userResult"`
	Choice     Choice `json:"choice" gorm:"foreignKey:UserResult;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// LongAnswerSubmission struct to store the LongAnswer Submission type data
type LongAnswerSubmission struct {
	LongAnswer   LongAnswer `json:"longAnswer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LongAnswerID uint       `json:"longAnswerID" gorm:"primary_key;autoIncrement:false"`
	Submission   `json:"submission"`
	UserResult   string `json:"userResult"`
}

// AfterSave updates the quiz grade when the user change the submission
func (submission *MCQSubmission) AfterSave(tx *gorm.DB) (err error) {
	submissionsHook.UpdateQuizGrade(submission.MCQID, submission.UserID, tx)
	return
}
