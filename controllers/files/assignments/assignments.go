package assignments

import (
	"backend/aws"
	filesUtils "backend/controllers/files"
	awsDiagnostics "backend/diagnostics/aws"
	assignmentsModel "backend/models/assignments"
	classesModel "backend/models/organizations"
	"bytes"
	"fmt"
	"github.com/labstack/echo"
	"io"
)

// UploadQuestionsFile uploads the questions file of the assignment
func UploadQuestionsFile(c echo.Context, assignment *assignmentsModel.Assignment, class *classesModel.Class) (string, error) {
	return uploadAssignmentFiles(c, assignment, "questionsFile", class)
}

// UploadModelAnswerFile uploads the model answer file of the assignment
func UploadModelAnswerFile(c echo.Context, assignment *assignmentsModel.Assignment, class *classesModel.Class) (string, error) {
	return uploadAssignmentFiles(c, assignment, "modelAnswerFile", class)
}

func uploadAssignmentFiles(c echo.Context, assignment *assignmentsModel.Assignment, formFileName string, class *classesModel.Class) (string, error) {
	// read file from source
	fileName, src, err := filesUtils.ReadFromSource(c, formFileName)
	if err != nil {
		if err.Error() == "http: no such file" {
			return "", nil
		}
		return "", err
	}
	buffer := bytes.NewBuffer(nil)
	_, err = io.Copy(buffer, src)
	if err != nil {
		return "", err
	}
	err = deleteAssignmentFiles(assignment, formFileName, class.Organization.S3BucketName)
	if err != nil {
		return "", err
	}
	filePath := fmt.Sprintf("%s/assignments/%s/%s/%s",
		assignment.ClassHash, assignment.Hash, formFileName, fileName)
	fileLocation, err := aws.Upload(buffer, filePath, class.Organization.S3BucketName, class.Organization.CloudfrontDomain)
	if err != nil {
		awsDiagnostics.WriteAWSPartErr(err, "Upload Video Part")
		return "", err
	}
	return fileLocation, nil
}

func deleteAssignmentFiles(assignment *assignmentsModel.Assignment, formFileName, s3BucketName string) error {
	folderPath := fmt.Sprintf("%s/assignments/%s/%s",
		assignment.ClassHash, assignment.Hash, formFileName)
	return aws.DeleteFolder(folderPath, s3BucketName)
}