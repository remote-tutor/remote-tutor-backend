package quizzes

import (
	"gorm.io/gorm"
)

// UpdateQuizTotalMark updates the quiz total marks by calculating the SUM of the questions marks
func UpdateQuizTotalMark(quizID uint, tx *gorm.DB) {
	subQuery := tx.Select("SUM(total_mark)").Where("quiz_id = ?", quizID).Table("mcqs")
	tx.Table("quizzes").Where("id = ?", quizID).Update("total_mark", subQuery)
}
