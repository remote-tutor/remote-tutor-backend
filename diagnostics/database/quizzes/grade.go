package quizzes

import (
	"backend/diagnostics"
	gradesModel "backend/models/quizzes"
)

func WriteGradeErr(err error, errorType string, quizGrade *gradesModel.QuizGrade) {
	filepath := "database/quizzes/choices.log"
	diagnostics.WriteToSpecificDiagnosticsFile(filepath, errorType, err, quizGrade)
}
