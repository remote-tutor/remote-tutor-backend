package assignments

import (
	dbInstance "backend/database"
	submissionsModel "backend/models/assignments"
)

func CreateSubmission(submission *submissionsModel.Submission) {
	dbInstance.GetDBConnection().Create(submission)
}

func UpdateSubmission(submission *submissionsModel.Submission) {
	dbInstance.GetDBConnection().Save(submission)
}

func GetSubmissionByID(id uint) submissionsModel.Submission {
	var submission submissionsModel.Submission
	dbInstance.GetDBConnection().First(&submission, id)
	return submission
}