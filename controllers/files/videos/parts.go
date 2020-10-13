package videos

import (
	"backend/aws"
	filesUtils "backend/controllers/files"
	videoParts "backend/models/videos"
	"bytes"
	"fmt"
	"github.com/labstack/echo"
	"io"
	"os"
)

func UploadVideoPart(c echo.Context, videoID uint) (string, error) {
	fileName, src, err := filesUtils.ReadFromSource(c, "videoPart")
	if err != nil {
		writeError("UploadVideoPart Line 16", err)
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
		writeError("UploadVideoPart Line 27", err)
		return "", err
	}
	return fileLocation, nil
}

func DeleteVideoPart(part *videoParts.VideoPart) error {
	filePath := fmt.Sprintf("%d/%s", part.VideoID, part.Name)
	return aws.Delete(filePath)
}

func DeleteVideo(video *videoParts.Video, parts []videoParts.VideoPart) error {
	var err error
	for _, part := range parts {
		err = DeleteVideoPart(&part)
		if err != nil {
			return err
		}
	}
	folderPath := fmt.Sprintf("%d", video.ID)
	return aws.Delete(folderPath)
}

func writeError(val string, err error) {
	file, fileErr := os.OpenFile("aws.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if fileErr == nil {
		file.Write([]byte(val + "\n"))
		file.Write([]byte("Error:\t" + err.Error() + "\n"))
	}
}
