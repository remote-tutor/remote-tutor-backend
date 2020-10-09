package quizzes

import (
	dbInstance "backend/database"
	quizzesModel "backend/models/quizzes"
)

// CreateChoice inserts a new choice to the database
func CreateChoice(choice *quizzesModel.Choice) error {
	err := dbInstance.GetDBConnection().Create(choice).Error
	return err
}

// UpdateChoice updates the given choice in the database
func UpdateChoice(choice *quizzesModel.Choice) error {
	err := dbInstance.GetDBConnection().Save(choice).Error
	return err
}

// DeleteChoice deletes the given choice in the database
func DeleteChoice(choice *quizzesModel.Choice) error {
	err := dbInstance.GetDBConnection().Unscoped().Delete(choice).Error
	return err
}

// GetChoiceByID returns the choice with the specific ID
func GetChoiceByID(id uint) quizzesModel.Choice {
	var choice quizzesModel.Choice
	dbInstance.GetDBConnection().First(&choice, id)
	return choice
}
