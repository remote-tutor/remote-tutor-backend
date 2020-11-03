package users

import (
	"backend/diagnostics"
	usersModel "backend/models/users"
)

func WriteUserErr(err error, errorType string, user *usersModel.User) {
	filePath := "database/users/users.log"
	diagnostics.WriteToSpecificDiagnosticsFile(filePath, errorType, err, user)
}
