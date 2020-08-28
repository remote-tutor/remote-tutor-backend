package quizzes

import (
	dbInstance "backend/database"
	quizzesModel "backend/models/quizzes"
)

// CreateQuiz inserts a new quiz to the database
func CreateQuiz(quiz *quizzesModel.Quiz) {
	dbInstance.GetDBConnection().Create(quiz)
}
