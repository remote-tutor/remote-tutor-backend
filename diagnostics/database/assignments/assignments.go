package assignments

import (
	"backend/diagnostics"
	assignmentsModel "backend/models/assignments"
)

func WriteAssignmentErr(err error, errorType string, assignment *assignmentsModel.Assignment) {
	filepath := "database/assignments/assignments.log"
	diagnostics.WriteToSpecificDiagnosticsFile(filepath, errorType, err, assignment)
}
