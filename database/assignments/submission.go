package assignments

import (
	dbInstance "backend/database"
	submissionsModel "backend/models/assignments"
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

func DeleteSubmission(submission *submissionsModel.AssignmentSubmission) {
	dbInstance.GetDBConnection().Delete(submission)
}
