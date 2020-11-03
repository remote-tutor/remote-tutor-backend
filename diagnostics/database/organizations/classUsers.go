package organizations

import (
	"backend/diagnostics"
	classUsersModel "backend/models/organizations"
)

func WriteClassUserErr(err error, errorType string, classUser *classUsersModel.ClassUser) {
	filepath := "database/organizations/class-users.log"
	diagnostics.WriteToSpecificDiagnosticsFile(filepath, errorType, err, classUser)
}

