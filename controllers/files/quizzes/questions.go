package quizzes

import (
	"backend/aws"
	filesUtils "backend/controllers/files"
	awsDiagnostics "backend/diagnostics/aws"
	classesModel "backend/models/organizations"
	questionsModel "backend/models/quizzes"
	"bytes"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/labstack/echo"
	"io"
	"math/rand"
	"strings"
)

func UploadQuestionImage(c *echo.Context, mcq *questionsModel.MCQ, class *classesModel.Class) (string, error) {
	_, src, err := filesUtils.ReadFromSource(*c, "image")
	if err != nil {
		if err.Error() == "http: no such file" {
			return mcq.ImagePath, nil
		}
		return "", err
	}
	err = DeleteQuestionImage(mcq, class) // delete the old image before uploading new one
	if err != nil {
		return "", err
	}
	buffer := bytes.NewBuffer(nil)
	_, err = io.Copy(buffer, src)
	if err != nil {
		return "", err
	}
	mime := mimetype.Detect(buffer.Bytes())
	randomInt := rand.Intn(1000)
	filePath := fmt.Sprintf("%s/quizzes/%s/%s-%d%s", class.Hash, mcq.Quiz.Hash, mcq.Hash, randomInt, mime.Extension())

	fileLocation, err := aws.Upload(buffer, filePath, &class.Organization)
	if err != nil {
		awsDiagnostics.WriteAWSUploadError(err, "Upload Question Image")
		return "", err
	}
	return fileLocation, nil
}

func DeleteQuestionImage(mcq *questionsModel.MCQ, class *classesModel.Class) error {
	if len(mcq.ImagePath) > 0 {
		filepath := strings.Split(mcq.ImagePath, ".net")[1]
		return aws.Delete(filepath, class.Organization.S3BucketName)
	}
	return nil
}

func DeleteQuizFiles(quiz *questionsModel.Quiz, class *classesModel.Class) error {
	folderPath := fmt.Sprintf("%s/quizzes/%s", class.Hash, quiz.Hash)
	return aws.DeleteFolder(folderPath, class.Organization.S3BucketName)
}