package quizzes

import (
	dbInstance "backend/database"
	"backend/database/diagnostics"
	quizzesModel "backend/models/quizzes"
)

// CreateMCQ inserts a new mcq question to the database
func CreateMCQ(mcq *quizzesModel.MCQ) error {
	err := dbInstance.GetDBConnection().Create(mcq).Error
	diagnostics.WriteError(err, "CreateMCQ")
	return err
}

// CreateLongAnswer inserts a new longAnswer question to the database
func CreateLongAnswer(longAnswer *quizzesModel.LongAnswer) {
	dbInstance.GetDBConnection().Create(longAnswer)
}

// UpdateMCQ updates mcq question in the database
func UpdateMCQ(mcq *quizzesModel.MCQ) error {
	err := dbInstance.GetDBConnection().Save(mcq).Error
	diagnostics.WriteError(err, "UpdateMCQ")
	return err
}

// UpdateLongAnswer updates long answer question in the database
func UpdateLongAnswer(longAnswer *quizzesModel.LongAnswer) {
	dbInstance.GetDBConnection().Save(longAnswer)
}

// DeleteMCQ deletes mcq question in the database
func DeleteMCQ(mcq *quizzesModel.MCQ) error {
	err := dbInstance.GetDBConnection().Unscoped().Delete(mcq).Error
	diagnostics.WriteError(err, "DeleteMCQ")
	return err
}

// DeleteLongAnswer deletes long answer question in the database
func DeleteLongAnswer(longAnswer *quizzesModel.LongAnswer) {
	dbInstance.GetDBConnection().Unscoped().Delete(longAnswer)
}

func GetMCQByID(id uint) quizzesModel.MCQ {
	var mcq quizzesModel.MCQ
	dbInstance.GetDBConnection().First(&mcq, id)
	return mcq
}

func GetLongAnswerByID(id uint) quizzesModel.LongAnswer {
	var longAnswer quizzesModel.LongAnswer
	dbInstance.GetDBConnection().First(&longAnswer, id)
	return longAnswer
}

// GetMCQsByQuiz retrievs all mcq questions for a quiz
func GetMCQsByQuiz(quizID uint) []quizzesModel.MCQ {
	mcqs := make([]quizzesModel.MCQ, 0)
	dbInstance.GetDBConnection().Preload("Choices").Where("quiz_id = ?", quizID).Find(&mcqs)
	return mcqs
}

// GetLongAnswersByQuiz retrievs all long answer questions for a quiz
func GetLongAnswersByQuiz(quizID uint) []quizzesModel.LongAnswer {
	longAnswers := make([]quizzesModel.LongAnswer, 0)
	dbInstance.GetDBConnection().Where("quiz_id = ?", quizID).Find(&longAnswers)
	return longAnswers
}
