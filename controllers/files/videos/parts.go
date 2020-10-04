package videos

import (
	"backend/aws"
	filesUtils "backend/controllers/files"
	"bytes"
	"fmt"
	"github.com/labstack/echo"
	"io"
)

func UploadVideoPart(c echo.Context, videoID uint) (string, error) {
	fileName, src, err := filesUtils.ReadFromSource(c, "videoPart")
	if err != nil {
		return "", err
	}
	buffer := bytes.NewBuffer(nil)
	_, err = io.Copy(buffer, src)
	if err != nil {
		return "", err
	}
	filePath := fmt.Sprintf("%d/%s", videoID, fileName)
	fileLocation, err := aws.Upload(buffer, filePath)
	if err != nil {
		return "", err
	}
	return fileLocation, nil
}