package quizzes

import (
	dbInstance "backend/database"
	quizzesModel "backend/models/quizzes"
)

// CreateMCQSubmission inserts a new mcq submission into the database
func CreateMCQSubmission(mcqSubmission *quizzesModel.MCQSubmission) error {
	err := dbInstance.GetDBConnection().Create(mcqSubmission).Error
	return err
}

// UpdateMCQSubmission updates an existing mcq submission in the database
func UpdateMCQSubmission(mcqSubmission *quizzesModel.MCQSubmission) error {
	err := dbInstance.GetDBConnection().Save(mcqSubmission).Error
	return err
}

// GetMCQSubmissionByQuestionAndUser retrieves the submission for a specific user for a specific question
func GetMCQSubmissionByQuestionAndUser(userID uint, mcqID uint) quizzesModel.MCQSubmission {
	var mcqSubmission quizzesModel.MCQSubmission
	dbInstance.GetDBConnection().Where("user_id = ? AND mcq_id = ?", userID, mcqID).Find(&mcqSubmission)
	return mcqSubmission
}

// GetMCQSubmissionsByQuestionID returns all the submissions of a specific question
func GetMCQSubmissionsByQuestionID(mcqID uint) []quizzesModel.MCQSubmission {
	var mcqSubmission []quizzesModel.MCQSubmission
	dbInstance.GetDBConnection().Where("mcq_id = ?", mcqID).Find(&mcqSubmission)
	return mcqSubmission
}

// GetMCQSubmissionsByQuizID retrieves all the mcq submissions for a specific user for a specific quiz
func GetMCQSubmissionsByQuizID(userID uint, quizID uint) []quizzesModel.MCQSubmission {
	db := dbInstance.GetDBConnection()
	mcqSubmissions := make([]quizzesModel.MCQSubmission, 0)
	subQuery := db.Table("mcqs").Select("id").Where("quiz_id = ? AND user_id = ?", quizID, userID)
	db.Where("mcq_id IN (?)", subQuery).Find(&mcqSubmissions)
	return mcqSubmissions
}

// GetLongAnswerSubmissionsByQuizID retrieves all the longAnswer submissions for a specific quiz
func GetLongAnswerSubmissionsByQuizID(userID uint, quizID uint) []quizzesModel.LongAnswerSubmission {
	db := dbInstance.GetDBConnection()
	longAnswerSubmission := make([]quizzesModel.LongAnswerSubmission, 0)
	subQuery := db.Table("mcqs").Select("id").Where("quiz_id = ? AND user_id = ?", quizID, userID)
	db.Where("long_answer_id IN (?)", subQuery).Find(&longAnswerSubmission)
	return longAnswerSubmission
}
