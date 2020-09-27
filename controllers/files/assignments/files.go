package assignments

import (
	assignmentsModel "backend/models/assignments"
	"fmt"
	"github.com/labstack/echo"
	"io"
	"mime/multipart"
	"os"
)

// UploadQuestionsFile uploads the questions file of the assignment
func UploadQuestionsFile(c echo.Context, assignment *assignmentsModel.Assignment) (string, error) {
	return uploadAssignmentFiles(c, assignment, "questions")
}

// UploadModelAnswerFile uploads the model answer file of the assignment
func UploadModelAnswerFile(c echo.Context, assignment *assignmentsModel.Assignment) (string, error) {
	return uploadAssignmentFiles(c, assignment, "modelAnswer")
}

func uploadAssignmentFiles(c echo.Context, assignment *assignmentsModel.Assignment, formFileName string) (string, error) {
	// read file from source
	fileName, src, err := readFromSource(c, formFileName)
	if err != nil {
		if err.Error() == "http: no such file" {
			return "", nil
		}
		return "", err
	}
	folderPath := fmt.Sprintf("assignmentsFiles/assignment %d", assignment.ID)
	createDirectoryIfNotExist(folderPath)
	fullFileName := fmt.Sprintf("%s/%s %s", folderPath, formFileName, fileName)
	err = copyFileToDestination(fullFileName, src)
	if err != nil {
		return "", err
	}
	return fullFileName, nil
}

func readFromSource(c echo.Context, formFileName string) (string, multipart.File, error) {
	file, err := c.FormFile(formFileName)
	if err != nil {
		return "", nil, err
	}
	src, err := file.Open()
	if err != nil {
		return "", nil, err
	}
	defer src.Close()
	return file.Filename, src, nil
}

func copyFileToDestination(fullFileName string, src multipart.File) error {
	// add file to destination
	dst, err := os.Create(fullFileName)
	if err != nil {
		return err
	}
	// copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

func createDirectoryIfNotExist(directoryName string) {
	if _, err := os.Stat(directoryName); os.IsNotExist(err) {
		os.Mkdir(directoryName, os.FileMode(int(0777)))
	}
}
