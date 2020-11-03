package quizzes

import (
	"backend/diagnostics"
	choiceModel "backend/models/quizzes"
)

func WriteChoiceErr(err error, errorType string, choice *choiceModel.Choice) {
	filepath := "database/quizzes/choices.log"
	diagnostics.WriteToSpecificDiagnosticsFile(filepath, errorType, err, choice)
}
