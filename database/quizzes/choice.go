package quizzes

import (
	dbInstance "backend/database"
	"backend/database/diagnostics"
	quizzesModel "backend/models/quizzes"
)

// CreateChoice inserts a new choice to the database
func CreateChoice(choice *quizzesModel.Choice) error {
	err := dbInstance.GetDBConnection().Create(choice).Error
	diagnostics.WriteError(err, "CreateChoice")
	return err
}

// UpdateChoice updates the given choice in the database
func UpdateChoice(choice *quizzesModel.Choice) error {
	err := dbInstance.GetDBConnection().Save(choice).Error
	diagnostics.WriteError(err, "UpdateChoice")
	return err
}

// DeleteChoice deletes the given choice in the database
func DeleteChoice(choice *quizzesModel.Choice) error {
	err := dbInstance.GetDBConnection().Unscoped().Delete(choice).Error
	diagnostics.WriteError(err, "DeleteChoice")
	return err
}

// GetChoiceByID returns the choice with the specific ID
func GetChoiceByID(id uint) quizzesModel.Choice {
	var choice quizzesModel.Choice
	dbInstance.GetDBConnection().First(&choice, id)
	return choice
}
