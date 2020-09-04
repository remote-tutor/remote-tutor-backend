package quizzes

import (
	dbInstance "backend/database"
	quizzesModel "backend/models/quizzes"
)

// CreateMCQSubmission inserts a new mcq submission into the database
func CreateMCQSubmission(mcqSubmission *quizzesModel.MCQSubmission) {
	dbInstance.GetDBConnection().Create(mcqSubmission)
}

// UpdateMCQSubmission updates an existing mcq submission in the database
func UpdateMCQSubmission(mcqSubmission *quizzesModel.MCQSubmission) {
	dbInstance.GetDBConnection().Save(mcqSubmission)
}

// GetMCQSubmissionByQuestionID retrieves the submission for a specific user for a specific question
func GetMCQSubmissionByQuestionID(userID uint, mcqID uint) quizzesModel.MCQSubmission {
	var mcqSubmission quizzesModel.MCQSubmission
	dbInstance.GetDBConnection().Where("user_id = ? AND question_id = ?", userID, mcqID).Find(&mcqSubmission)
	return mcqSubmission
}

// GetMCQSubmissionsByQuizID retrieves all the mcq submissions for a specific user for a specific quiz
func GetMCQSubmissionsByQuizID(userID uint, quizID uint) []quizzesModel.MCQSubmission {
	db := dbInstance.GetDBConnection()
	var mcqSubmissions []quizzesModel.MCQSubmission
	subQuery := db.Table("mcqs").Select("id").Where("quiz_id = ? AND user_id = ?", quizID, userID)
	db.Where("question_id IN (?)", subQuery).Find(&mcqSubmissions)
	return mcqSubmissions
}

// GetLongAnswerSubmissionsByQuizID retrieves all the longAnswer submissions for a specific quiz
func GetLongAnswerSubmissionsByQuizID(userID uint, quizID uint) []quizzesModel.LongAnswerSubmission {
	db := dbInstance.GetDBConnection()
	var longAnswerSubmission []quizzesModel.LongAnswerSubmission
	subQuery := db.Table("mcqs").Select("id").Where("quiz_id = ? AND user_id = ?", quizID, userID)
	db.Where("question_id IN (?)", subQuery).Find(&longAnswerSubmission)
	return longAnswerSubmission
}
