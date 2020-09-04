package quizzes

import (
	"gorm.io/gorm"
)

// Question struct to store the question data
type Question struct {
	gorm.Model
	Text      string `json:"text"`
	TotalMark int    `json:"totalMark"`
	Quiz      Quiz   `json:"quiz"`
	QuizID    uint   `json:"quizID"`
}

// MCQ struct to store the MCQ question type data
type MCQ struct {
	Question      `json:"question"`
	CorrectAnswer uint     `json:"correctAnswer"`
	Choices       []Choice `json:"choices"`
}

// LongAnswer struct to store the LongAnswer question type data
type LongAnswer struct {
	Question      `json:"question"`
	CorrectAnswer string `json:"correctAnswer"`
}
