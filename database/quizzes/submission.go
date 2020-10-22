package quizzes

import (
	dbInstance "backend/database"
	submissionDiagnostics "backend/diagnostics/database/quizzes"
	quizzesModel "backend/models/quizzes"
	"gorm.io/gorm/clause"
)


// CreateOrUpdateMCQSubmission creates OR (updates on conflict) mcq submission in the database
func CreateOrUpdateMCQSubmission(mcqSubmission *quizzesModel.MCQSubmission) error {
	err := dbInstance.GetDBConnection().Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "user_id"}, {Name: "mcq_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"user_result", "grade", "updated_at"}),
	}).Create(mcqSubmission).Error
	submissionDiagnostics.WriteSubmissionErr(err, "CreateOrUpdate", mcqSubmission)
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
