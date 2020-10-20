package quizzes

import (
	"backend/diagnostics"
	quizzesModel "backend/models/quizzes"
)

func WriteQuizErr(err error, errorType string, quiz *quizzesModel.Quiz) {
	filepath := "database/quizzes/quizzes.log"
	diagnostics.WriteToSpecificDiagnosticsFile(filepath, errorType, err, quiz)
}
