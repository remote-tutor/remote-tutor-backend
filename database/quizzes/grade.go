package quizzes

import (
	dbInstance "backend/database"
	quizzesModel "backend/models/quizzes"
)

// CreateGrade inserts a new quiz grade to the database
func CreateGrade(grade *quizzesModel.QuizGrade) {
	dbInstance.GetDBConnection().Create(grade)
}

// UpdateGrade updates an existing quiz grade in the database
func UpdateGrade(grade *quizzesModel.QuizGrade) {
	dbInstance.GetDBConnection().Save(grade)
}

// GetGradesByQuizID gets grade for logged-in user for a specific quiz
func GetGradesByQuizID(userID uint, quizID uint) quizzesModel.QuizGrade {
	var quizGrade quizzesModel.QuizGrade
	dbInstance.GetDBConnection().Where("user_id = ? AND quiz_id = ?", userID, quizID).FirstOrInit(&quizGrade)
	return quizGrade
}

// GetGradesByQuizForAllUsers retrieves all class grades for a specific quiz
func GetGradesByQuizForAllUsers(quizID uint) []quizzesModel.QuizGrade {
	var quizGrades []quizzesModel.QuizGrade
	dbInstance.GetDBConnection().Where("quiz_id = ?", quizID).Joins("User").Order("full_name").Find(&quizGrades)
	return quizGrades
}
