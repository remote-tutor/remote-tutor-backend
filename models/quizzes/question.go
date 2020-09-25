package quizzes

import (
	quizzesHooks "backend/hooks/quizzes"

	"gorm.io/gorm"
)

// Question struct to store the question data
type Question struct {
	ID        uint   `json:"ID"`
	Text      string `json:"text"`
	TotalMark int    `json:"totalMark"`
	Quiz      Quiz   `json:"quiz" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	QuizID    uint   `json:"quizID"`
	ImagePath string `json:"imagePath"`
	Image     []byte `json:"image" gorm:"-"`
}

// MCQ struct to store the MCQ question type data
type MCQ struct {
	Question      `json:"question"`
	CorrectAnswer uint     `json:"correctAnswer"`
	Choices       []Choice `json:"choices" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// LongAnswer struct to store the LongAnswer question type data
type LongAnswer struct {
	Question      `json:"question"`
	CorrectAnswer string `json:"correctAnswer"`
}

// AfterSave updates the quiz total mark every time question is created or updated
func (mcq *MCQ) AfterSave(tx *gorm.DB) (err error) {
	quizzesHooks.UpdateQuizTotalMark(mcq.QuizID, tx)
	return
}

// AfterUpdate updates the user's submissions after question is updated
// func (mcq *MCQ) AfterUpdate(tx *gorm.DB) (err error) {
// 	quizzesHooks.UpdateQuizTotalMark(mcq.QuizID, tx)
// 	return
// }

// AfterDelete updates the quiz total mark every time question is deleted
func (mcq *MCQ) AfterDelete(tx *gorm.DB) (err error) {
	quizzesHooks.UpdateQuizTotalMark(mcq.QuizID, tx)
	quizzesHooks.UpdateQuizGradeForAllUsers(mcq.QuizID, tx)
	return
}
