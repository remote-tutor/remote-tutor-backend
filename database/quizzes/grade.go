package quizzes

import (
	dbInstance "backend/database"
	gradesDiagnostics "backend/diagnostics/database/quizzes"
	quizzesModel "backend/models/quizzes"
	"errors"
	"gorm.io/gorm"
	"time"
)

// GetGradesByQuizID gets grade for logged-in user for a specific quiz
func GetGradesByQuizID(userID uint, quizID uint) quizzesModel.QuizGrade {
	var quizGrade quizzesModel.QuizGrade
	dbInstance.GetDBConnection().Where("user_id = ? AND quiz_id = ?", userID, quizID).Preload("User").FirstOrInit(&quizGrade)
	return quizGrade
}

// CreateGrade inserts a new quiz grade to the database
func CreateGrade(grade *quizzesModel.QuizGrade) error {
	err := dbInstance.GetDBConnection().FirstOrCreate(grade).Error
	gradesDiagnostics.WriteGradeErr(err, "Create", grade)
	return err
}

// GetStudentRemainingTime fetches the remaining time for a quiz by given student
func GetStudentRemainingTime(userID, quizID uint) (time.Time, bool) {
	var quizGrade quizzesModel.QuizGrade
	err := dbInstance.GetDBConnection().Where("user_id = ? AND quiz_id = ?", userID, quizID).First(&quizGrade).Error
	recordNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	return quizGrade.ValidTill, !recordNotFound
}

// GetGradesByMonthForAllUsers gets the grades of all quizzes in a specific month
func GetGradesByMonthForAllUsers(class string, startOfMonth, endOfMonth time.Time) ([]quizzesModel.Quiz, []quizzesModel.QuizGrade) {
	quizzes := GetQuizzesByClassAndMonthAndYear(class, startOfMonth, endOfMonth)
	quizzesIDs := make([]uint, len(quizzes))
	for i := 0; i < len(quizzes); i++ {
		quizzesIDs[i] = quizzes[i].ID
	}
	quizzesGrades := make([]quizzesModel.QuizGrade, 0)
	dbInstance.GetDBConnection().Where("quiz_id IN (?)", quizzesIDs).
		Joins("User").Order("full_name").Find(&quizzesGrades)
	return quizzes, quizzesGrades
}

func GetGradesByMonthAndUser(class string, userID uint, startOfMonth, endOfMonth time.Time) ([]quizzesModel.Quiz, []quizzesModel.QuizGrade) {
	quizzes := GetQuizzesByClassAndMonthAndYear(class, startOfMonth, endOfMonth)
	quizzesIDs := make([]uint, len(quizzes))
	for i := 0; i < len(quizzes); i++ {
		quizzesIDs[i] = quizzes[i].ID
	}
	quizzesGrades := make([]quizzesModel.QuizGrade, 0)
	dbInstance.GetDBConnection().Where("quiz_id IN (?) AND user_id = ?", quizzesIDs, userID).
		Joins("User").Order("full_name").Find(&quizzesGrades)
	return quizzes, quizzesGrades
}
