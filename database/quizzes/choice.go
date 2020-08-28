package quizzes

import (
	dbInstance "backend/database"
	quizzesModel "backend/models/quizzes"
)

// CreateChoice inserts a new choice to the database
func CreateChoice(choice *quizzesModel.Choice) {
	dbInstance.GetDBConnection().Create(choice)
}
