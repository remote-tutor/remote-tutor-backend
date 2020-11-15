package quizzes

import (
	dbInstance "backend/database"
	gradesDiagnostics "backend/diagnostics/database/quizzes"
	quizzesModel "backend/models/quizzes"
	"errors"
	"gorm.io/gorm"
	"time"
)

// CreateGrade inserts a new quiz grade to the database
func CreateGrade(grade *quizzesModel.QuizGrade) error {
	err := dbInstance.GetDBConnection().FirstOrCreate(grade).Error
	gradesDiagnostics.WriteGradeErr(err, "Create", grade)
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

// GetStudentRemainingTime fetches the remaining time for a quiz by given student
func GetStudentRemainingTime(userID, quizID uint) (time.Time, bool) {
	var quizGrade quizzesModel.QuizGrade
	err := dbInstance.GetDBConnection().Where("user_id = ? AND quiz_id = ?", userID, quizID).First(&quizGrade).Error
	recordNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	return quizGrade.ValidTill, !recordNotFound
}

// GetGradesByQuizForAllUsers retrieves all class grades for a specific quiz
func GetGradesByQuizForAllUsers(quizID uint) []quizzesModel.QuizGrade {
	quizGrades := make([]quizzesModel.QuizGrade, 0)
	dbInstance.GetDBConnection().Where("quiz_id = ?", quizID).Joins("User").Order("full_name").Find(&quizGrades)
	return quizGrades
}
