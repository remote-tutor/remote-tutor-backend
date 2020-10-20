package quizzes

import (
	"backend/diagnostics"
	submissionsModel "backend/models/quizzes"
)

func WriteSubmissionErr(err error, errorType string, submission *submissionsModel.MCQSubmission) {
	filePath := "database/quizzes/submissions.log"
	diagnostics.WriteToSpecificDiagnosticsFile(filePath, errorType, err, submission)
}
