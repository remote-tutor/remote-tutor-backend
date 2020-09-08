package quizzes

import (
	dbInstance "backend/database"
	quizzesModel "backend/models/quizzes"
	"time"
)

// CreateQuiz inserts a new quiz to the database
func CreateQuiz(quiz *quizzesModel.Quiz) {
	dbInstance.GetDBConnection().Create(quiz)
}

//GetPastQuizzes retrieves list of past quizzes
func GetPastQuizzes(year int) []quizzesModel.Quiz {
	pastQuizzes := make([]quizzesModel.Quiz, 0)
	dbInstance.GetDBConnection().Where("year = ? AND end_time < ?", year, time.Now()).Find(&pastQuizzes)
	return pastQuizzes
}

//GetFutureQuizzes retrieves list of future quizzes
func GetFutureQuizzes(year int) []quizzesModel.Quiz {
	futureQuizzes := make([]quizzesModel.Quiz, 0)
	dbInstance.GetDBConnection().Where("year = ? AND start_time > ?", year, time.Now()).Find(&futureQuizzes)
	return futureQuizzes
}

//GetCurrentQuizzes retrieves list of current quizzes
func GetCurrentQuizzes(year int) []quizzesModel.Quiz {
	currentQuizzes := make([]quizzesModel.Quiz, 0)
	currentTime := time.Now()
	dbInstance.GetDBConnection().Where("year = ? AND start_time < ? AND end_time > ?", year, currentTime, currentTime).Find(&currentQuizzes)
	return currentQuizzes
}

// DeleteQuiz deletes the specified quiz from the database
func DeleteQuiz(quiz *quizzesModel.Quiz) {
	dbInstance.GetDBConnection().Unscoped().Delete(quiz)
}

// UpdateQuiz updates the quiz in the database
func UpdateQuiz(quiz *quizzesModel.Quiz) {
	dbInstance.GetDBConnection().Save(quiz)
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
