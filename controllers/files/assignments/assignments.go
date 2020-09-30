package assignments

import (
	filesUtils "backend/controllers/files"
	assignmentsModel "backend/models/assignments"
	"fmt"
	"github.com/labstack/echo"
)

// UploadQuestionsFile uploads the questions file of the assignment
func UploadQuestionsFile(c echo.Context, assignment *assignmentsModel.Assignment) (string, error) {
	return uploadAssignmentFiles(c, assignment, "questionsFile")
}

// UploadModelAnswerFile uploads the model answer file of the assignment
func UploadModelAnswerFile(c echo.Context, assignment *assignmentsModel.Assignment) (string, error) {
	return uploadAssignmentFiles(c, assignment, "modelAnswerFile")
}

func uploadAssignmentFiles(c echo.Context, assignment *assignmentsModel.Assignment, formFileName string) (string, error) {
	// read file from source
	fileName, src, err := filesUtils.ReadFromSource(c, formFileName)
	if err != nil {
		if err.Error() == "http: no such file" {
			return "", nil
		}
		return "", err
	}
	assignmentFolderPath := fmt.Sprintf("assignmentsFiles/assignment %d", assignment.ID)
	filesUtils.CreateDirectoryIfNotExist(assignmentFolderPath)

	subFolderPath := fmt.Sprintf("%s/%s", assignmentFolderPath, formFileName)
	filesUtils.DeleteDirectory(subFolderPath)
	filesUtils.CreateDirectoryIfNotExist(subFolderPath)

	fullFileName := fmt.Sprintf("%s/%s", subFolderPath, fileName)
	err = filesUtils.CopyFileToDestination(fullFileName, src)
	if err != nil {
		return "", err
	}
	return fullFileName, nil
}