package assignments

import (
	dbInstance "backend/database"
	dbPagination "backend/database/scopes"
	"backend/diagnostics"
	submissionsModel "backend/models/assignments"
	"fmt"
	"gorm.io/gorm"
)

func CreateSubmission(submission *submissionsModel.AssignmentSubmission) error {
	err := dbInstance.GetDBConnection().Create(submission).Error
	diagnostics.WriteError(err, "database.log", "CreateSubmission (assignment)")
	return err
}

func UpdateSubmission(submission *submissionsModel.AssignmentSubmission) error {
	err := dbInstance.GetDBConnection().Save(submission).Error
	diagnostics.WriteError(err, "database.log", "UpdateSubmission (assignment)")
	return err
}

func GetSubmissionByUserAndAssignment(userID, assignmentID uint) submissionsModel.AssignmentSubmission {
	var submission submissionsModel.AssignmentSubmission
	dbInstance.GetDBConnection().Where("user_id = ? AND assignment_id = ?", userID, assignmentID).Find(&submission)
	return submission
}

func GetSubmissionsByAssignmentForAllUsers(paginationData *dbPagination.PaginationData, assignmentID uint, fullNameSearch string) ([]submissionsModel.AssignmentSubmission, int64) {
	submissions := make([]submissionsModel.AssignmentSubmission, 0)
	db := dbInstance.GetDBConnection().
		Where("assignment_id = ? AND full_name LIKE ?", assignmentID, fmt.Sprintf("%%%s%%", fullNameSearch)).
		Joins("User")
	totalSubmissions := countSubmissions(db)
	db.Scopes(dbPagination.Paginate(paginationData)).Find(&submissions)
	return submissions, totalSubmissions
}

func countSubmissions(db *gorm.DB) int64 {
	totalSubmissions := int64(0)
	db.Model(&submissionsModel.AssignmentSubmission{}).Count(&totalSubmissions)
	return totalSubmissions
}

func DeleteSubmission(submission *submissionsModel.AssignmentSubmission) error {
	err := dbInstance.GetDBConnection().Delete(submission).Error
	diagnostics.WriteError(err, "database.log", "DeleteSubmission (assignment)")
	return err
}
