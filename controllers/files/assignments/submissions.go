package assignments

import (
	filesUtils "backend/controllers/files"
	assignmentsModel "backend/models/assignments"
	"fmt"
	"github.com/labstack/echo"
)

func UploadUserSubmissionFile(c echo.Context, submission *assignmentsModel.AssignmentSubmission) (string, error) {
	// read file from source
	fileName, src, err := filesUtils.ReadFromSource(c, "submissionFile")
	if err != nil {
		if err.Error() == "http: no such file" {
			return "", nil
		}
		return "", err
	}
	submissionsSubFolder := fmt.Sprintf("assignmentsFiles/assignment %d/submissions", submission.AssignmentID)
	filesUtils.CreateDirectoryIfNotExist(submissionsSubFolder)
	userSubmissionSubFolder := fmt.Sprintf("%s/user %d", submissionsSubFolder, submission.UserID)
	filesUtils.DeleteDirectory(userSubmissionSubFolder)
	filesUtils.CreateDirectoryIfNotExist(userSubmissionSubFolder)
	fullFileName := fmt.Sprintf("%s/%s", userSubmissionSubFolder, fileName)
	err = filesUtils.CopyFileToDestination(fullFileName, src)
	if err != nil {
		return "", err
	}
	return fullFileName, nil
}


