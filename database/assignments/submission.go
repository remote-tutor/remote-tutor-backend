package assignments

import (
	dbInstance "backend/database"
	dbPagination "backend/database/scopes"
	submissionsDiagnostics "backend/diagnostics/database/assignments"
	submissionsModel "backend/models/assignments"
	classesModel "backend/models/organizations"
	"fmt"
	"gorm.io/gorm"
)

func CreateSubmission(submission *submissionsModel.AssignmentSubmission) error {
	err := dbInstance.GetDBConnection().Create(submission).Error
	submissionsDiagnostics.WriteSubmissionErr(err, "Create", submission)
	return err
}

func UpdateSubmission(submission *submissionsModel.AssignmentSubmission) error {
	err := dbInstance.GetDBConnection().Save(submission).Error
	submissionsDiagnostics.WriteSubmissionErr(err, "Update", submission)
	return err
}

func GetSubmissionByUserAndAssignment(userID uint, assignmentHash string) submissionsModel.AssignmentSubmission {
	subQuery := dbInstance.GetDBConnection().Select("id").
		Where("assignments.hash = ?", assignmentHash).Table("assignments")
	var submission submissionsModel.AssignmentSubmission
	dbInstance.GetDBConnection().Where("user_id = ? AND assignment_id = (?)",
		userID, subQuery).Find(&submission)
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
	submissionsDiagnostics.WriteSubmissionErr(err, "Delete", submission)
	return err
}

func GetClassByAssignmentHash(assignmentHash string) classesModel.Class {
	var class classesModel.Class
	subQuery := dbInstance.GetDBConnection().Select("class_hash").
		Where("assignments.hash = ?", assignmentHash).Table("assignments")
	dbInstance.GetDBConnection().Where("classes.hash = (?)", subQuery).Preload("Organization").Find(&class)
	return class
}