package quizzes

import (
	dbInstance "backend/database"
	"backend/diagnostics"
	quizzesModel "backend/models/quizzes"
)

// CreateGrade inserts a new quiz grade to the database
func CreateGrade(grade *quizzesModel.QuizGrade) error {
	err := dbInstance.GetDBConnection().FirstOrCreate(grade).Error
	diagnostics.WriteError(err, "database.log", "CreateGrade")
	return err
}

// UpdateGrade updates an existing quiz grade in the database
func UpdateGrade(grade *quizzesModel.QuizGrade) {
	dbInstance.GetDBConnection().Save(grade)
}

// GetGradesByQuizID gets grade for logged-in user for a specific quiz
func GetGradesByQuizID(userID uint, quizID uint) quizzesModel.QuizGrade {
	var quizGrade quizzesModel.QuizGrade
	dbInstance.GetDBConnection().Where("user_id = ? AND quiz_id = ?", userID, quizID).Preload("User").FirstOrInit(&quizGrade)
	return quizGrade
}

// GetGradesByQuizForAllUsers retrieves all class grades for a specific quiz
func GetGradesByQuizForAllUsers(quizID uint) []quizzesModel.QuizGrade {
	quizGrades := make([]quizzesModel.QuizGrade, 0)
	dbInstance.GetDBConnection().Where("quiz_id = ?", quizID).Joins("User").Order("full_name").Find(&quizGrades)
	return quizGrades
}
