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

func GetSubmissionByID(id uint) submissionsModel.AssignmentSubmission {
	var submission submissionsModel.AssignmentSubmission
	dbInstance.GetDBConnection().First(&submission, id)
	return submission
}

func DeleteSubmission(submission *submissionsModel.AssignmentSubmission) {
	dbInstance.GetDBConnection().Delete(submission)
}