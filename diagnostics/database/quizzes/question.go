package quizzes

import (
	"backend/diagnostics"
	questionsModel "backend/models/quizzes"
)

func WriteQuestionErr(err error, errorType string, question *questionsModel.MCQ) {
	filePath := "database/quizzes/questions.log"
	diagnostics.WriteToSpecificDiagnosticsFile(filePath, errorType, err, question)
}
