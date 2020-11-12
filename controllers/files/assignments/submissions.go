package assignments

import (
	"backend/aws"
	filesUtils "backend/controllers/files"
	awsDiagnostics "backend/diagnostics/aws"
	classesModel "backend/models/organizations"
	"bytes"
	"fmt"
	"github.com/labstack/echo"
	"io"
)

func UploadUserSubmissionFile(c echo.Context, userID uint, assignmentHash string, class *classesModel.Class) (string, error) {
	// read file from source
	fileName, src, err := filesUtils.ReadFromSource(c, "submissionFile")
	if err != nil {
		if err.Error() == "http: no such file" {
			return "", nil
		}
		return "", err
	}
	defer src.Close()
	filePath := fmt.Sprintf("%s/assignments/%s/submissions/%d/%s",
		class.Hash, assignmentHash, userID, fileName)
	buffer := bytes.NewBuffer(nil)
	_, err = io.Copy(buffer, src)
	if err != nil {
		return "", err
	}
	err = deleteSubmissionFiles(class.Hash, class.Organization.S3BucketName, assignmentHash, userID)
	if err != nil {
		return "", err
	}
	fileLocation, err := aws.Upload(buffer, filePath, &class.Organization)
	if err != nil {
		awsDiagnostics.WriteAWSUploadError(err, "Upload Submission")
		return "", err
	}
	return fileLocation, nil
}

func deleteSubmissionFiles(classHash, s3BucketName, assignmentHash string, userID uint) error {
	folderPath := fmt.Sprintf("%s/assignments/%s/submissions/%d",
		classHash, assignmentHash, userID)
	return aws.DeleteFolder(folderPath, s3BucketName)
}
