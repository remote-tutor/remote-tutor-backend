package quizzes

import (
	dbInstance "backend/database"
	quizzesModel "backend/models/quizzes"
)

// CreateChoice inserts a new choice to the database
func CreateChoice(choice *quizzesModel.Choice) {
	dbInstance.GetDBConnection().Create(choice)
}

// UpdateChoice updates the given choice in the database
func UpdateChoice(choice *quizzesModel.Choice) {
	dbInstance.GetDBConnection().Save(choice)
}

// DeleteChoice deletes the given choice in the database
func DeleteChoice(choice *quizzesModel.Choice) {
	dbInstance.GetDBConnection().Unscoped().Delete(choice)
}

// GetChoiceByID returns the choice with the specific ID
func GetChoiceByID(id uint) quizzesModel.Choice {
	var choice quizzesModel.Choice
	dbInstance.GetDBConnection().First(&choice, id)
	return choice
}
