package quizzes

import (
	dbInstance "backend/database"
	dbPagination "backend/database/scopes"
	"backend/diagnostics"
	quizzesModel "backend/models/quizzes"
	"time"

	"gorm.io/gorm"
)

// CreateQuiz inserts a new quiz to the database
func CreateQuiz(quiz *quizzesModel.Quiz) error {
	err := dbInstance.GetDBConnection().Create(quiz).Error
	diagnostics.WriteError(err, "database.log", "CreateQuiz")
	return err
}

//GetPastQuizzes retrieves list of past quizzes
func GetPastQuizzes(paginationData *dbPagination.PaginationData, year int) ([]quizzesModel.Quiz, int64) {
	pastQuizzes := make([]quizzesModel.Quiz, 0)
	db := dbInstance.GetDBConnection().Where("year = ? AND end_time < ?", year, time.Now())
	totalQuizzes := countRequiredQuizzes(db)
	db.Scopes(dbPagination.Paginate(paginationData)).Find(&pastQuizzes)
	return pastQuizzes, totalQuizzes
}

//GetFutureQuizzes retrieves list of future quizzes
func GetFutureQuizzes(paginationData *dbPagination.PaginationData, year int) ([]quizzesModel.Quiz, int64) {
	futureQuizzes := make([]quizzesModel.Quiz, 0)
	db := dbInstance.GetDBConnection().Where("year = ? AND start_time > ?", year, time.Now())
	totalQuizzes := countRequiredQuizzes(db)
	db.Scopes(dbPagination.Paginate(paginationData)).Find(&futureQuizzes)
	return futureQuizzes, totalQuizzes
}

//GetCurrentQuizzes retrieves list of current quizzes
func GetCurrentQuizzes(paginationData *dbPagination.PaginationData, year int) ([]quizzesModel.Quiz, int64) {
	currentQuizzes := make([]quizzesModel.Quiz, 0)
	currentTime := time.Now()
	db := dbInstance.GetDBConnection().Where("year = ? AND start_time < ? AND end_time > ?", year, currentTime, currentTime)
	totalQuizzes := countRequiredQuizzes(db)
	db.Scopes(dbPagination.Paginate(paginationData)).Find(&currentQuizzes)
	return currentQuizzes, totalQuizzes
}

// countRequiredQuizzes gets the number of the quizzes that the user requests (current, future, or past)
func countRequiredQuizzes(db *gorm.DB) int64 {
	totalQuizzes := int64(0)
	db.Model(&quizzesModel.Quiz{}).Count(&totalQuizzes)
	return totalQuizzes
}

// DeleteQuiz deletes the specified quiz from the database
func DeleteQuiz(quiz *quizzesModel.Quiz) error {
	err := dbInstance.GetDBConnection().Unscoped().Delete(quiz).Error
	diagnostics.WriteError(err, "database.log", "DeleteQuiz")
	return err
}

// UpdateQuiz updates the quiz in the database
func UpdateQuiz(quiz *quizzesModel.Quiz) error {
	err := dbInstance.GetDBConnection().Save(quiz).Error
	diagnostics.WriteError(err, "database.log", "UpdateQuiz")
	return err
}

// GetQuizByID retrieves the quiz by the quizID
func GetQuizByID(id uint) quizzesModel.Quiz {
	var quiz quizzesModel.Quiz
	dbInstance.GetDBConnection().First(&quiz, id)
	return quiz
}

// GetQuizzesByMonthAndYear retrieves all the quizzes within a month period.
func GetQuizzesByMonthAndYear(year int, startOfMonth, endOfMonth time.Time) []quizzesModel.Quiz {
	quizzes := make([]quizzesModel.Quiz, 0)
	currentTime := time.Now()
	dbInstance.GetDBConnection().
		Where("year = ? AND start_time >= ? AND end_time <= ? AND end_time < ?",
			year, startOfMonth, endOfMonth, currentTime).
		Order("start_time").
		Find(&quizzes)
	return quizzes
}
