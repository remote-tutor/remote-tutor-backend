package assignments

import (
	dbInstance "backend/database"
	"backend/database/scopes"
	submissionsModel "backend/models/assignments"
	"fmt"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func CreateSubmission(submission *submissionsModel.AssignmentSubmission) {
	dbInstance.GetDBConnection().Create(submission)
}

func UpdateSubmission(submission *submissionsModel.AssignmentSubmission) {
	dbInstance.GetDBConnection().Save(submission)
}

func GetSubmissionByUserAndAssignment(userID, assignmentID uint) submissionsModel.AssignmentSubmission {
	var submission submissionsModel.AssignmentSubmission
	dbInstance.GetDBConnection().Where("user_id = ? AND assignment_id = ?", userID, assignmentID).Find(&submission)
	return submission
}

func GetSubmissionsByAssignmentForAllUsers(c echo.Context, assignmentID uint, fullNameSearch string) ([]submissionsModel.AssignmentSubmission, int64) {
	submissions := make([]submissionsModel.AssignmentSubmission, 0)
	db := dbInstance.GetDBConnection().
		Where("assignment_id = ? AND full_name LIKE ?", assignmentID, fmt.Sprintf("%%%s%%", fullNameSearch)).
		Joins("User")
	totalSubmissions := countSubmissions(db)
	db.Scopes(scopes.Paginate(c)).Find(&submissions)
	return submissions, totalSubmissions
}

func countSubmissions(db *gorm.DB) int64 {
	totalSubmissions := int64(0)
	db.Model(&submissionsModel.AssignmentSubmission{}).Count(&totalSubmissions)
	return totalSubmissions
}

func DeleteSubmission(submission *submissionsModel.AssignmentSubmission) {
	dbInstance.GetDBConnection().Delete(submission)
}
