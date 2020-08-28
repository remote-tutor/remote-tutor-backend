package quizzes

import (
	dbInstance "backend/database"
	quizzesModel "backend/models/quizzes"
)

// CreateMCQ inserts a new mcq question to the database
func CreateMCQ(mcq *quizzesModel.MCQ) {
	dbInstance.GetDBConnection().Create(mcq)
}

// CreateLongAnswer inserts a new longAnswer question to the database
func CreateLongAnswer(longAnswer *quizzesModel.LongAnswer) {
	dbInstance.GetDBConnection().Create(longAnswer)
}

// GetMCQQuestionsByQuiz retrievs all mcq questions for a quiz
func GetMCQQuestionsByQuiz(quizID uint) []quizzesModel.MCQ {
	var mcqs []quizzesModel.MCQ
	dbInstance.GetDBConnection().Where("quiz_id = ?", quizID).Find(&mcqs)
	return mcqs
}

// GetLongAnswerQuestionsByQuiz retrievs all long answer questions for a quiz
func GetLongAnswerQuestionsByQuiz(quizID uint) []quizzesModel.LongAnswer {
	var longAnswers []quizzesModel.LongAnswer
	dbInstance.GetDBConnection().Where("quiz_id = ?", quizID).Find(&longAnswers)
	return longAnswers
}
