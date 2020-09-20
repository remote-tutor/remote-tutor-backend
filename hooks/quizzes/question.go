package quizzes

import "gorm.io/gorm"

// UpdateQuizGradeForAllUsers updates quiz grades for all users
func UpdateQuizGradeForAllUsers(quizID uint, tx *gorm.DB) {
	mcqsIDsSubQuery := tx.Select("id").Where("quiz_id = (?)", quizID).Table("mcqs")
	var usersIDs []int
	tx.Distinct().Table("mcq_submissions").
		Select("user_id").Where("mcq_id IN (?)", mcqsIDsSubQuery).Scan(&usersIDs)
	for _, userID := range usersIDs {
		subQuery := tx.Select("SUM(grade)").Where("mcq_id IN (?) AND user_id = ?",
			mcqsIDsSubQuery, userID).Table("mcq_submissions")
		tx.Table("quiz_grades").Where("user_id = ? AND quiz_id = ?", userID, quizID).
			Update("grade", subQuery)
	}
}
