package assignments

import (
	"backend/diagnostics"
	submissionsModel "backend/models/assignments"
)

func WriteSubmissionErr(err error, errorType string, submission *submissionsModel.AssignmentSubmission) {
	filepath := "database/assignments/submissions.log"
	diagnostics.WriteToSpecificDiagnosticsFile(filepath, errorType, err, submission)
}
