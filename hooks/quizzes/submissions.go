package quizzes

import "gorm.io/gorm"

// UpdateQuizGrade updates the total marks for a quiz for a specific user
func UpdateQuizGrade(mcqID, userID uint, tx *gorm.DB) {
	var quizID int
	tx.Select("quiz_id").Where("id = ?", mcqID).Table("mcqs").Scan(&quizID)
	mcqsIDsSubQuery := tx.Select("id").Where("quiz_id = (?)", quizID).Table("mcqs")

	subQuery := tx.Select("SUM(grade)").Where("mcq_id IN (?) AND user_id = ?",
		mcqsIDsSubQuery, userID).Table("mcq_submissions")
	tx.Table("quiz_grades").Where("user_id = ? AND quiz_id = ?", userID, quizID).
		Update("grade", subQuery)
}
