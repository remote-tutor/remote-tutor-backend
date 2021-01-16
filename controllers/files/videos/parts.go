package videos

import (
	"backend/aws"
	filesUtils "backend/controllers/files"
	awsDiagnostics "backend/diagnostics/aws"
	classesModel "backend/models/organizations"
	videoParts "backend/models/videos"
	"bytes"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/labstack/echo"
	"io"
	"strings"
)

func UploadVideoPart(c echo.Context, video *videoParts.Video, class *classesModel.Class) (string, bool, error) {
	fileName, src, err := filesUtils.ReadFromSource(c, "videoPart")
	if src != nil {
		defer src.Close()
	}
	if err != nil {
		awsDiagnostics.WriteAWSPartErr(err, "Upload Video Part (ReadFromSource)")
		return "", false, err
	}
	buffer := bytes.NewBuffer(nil)
	_, err = io.Copy(buffer, src)
	if err != nil {
		awsDiagnostics.WriteAWSPartErr(err, "Upload Video Part (Copy)")
		return "", false, err
	}
	filePath := fmt.Sprintf("%s/videos/%s/%s", video.ClassHash, video.Hash, fileName)
	mime := mimetype.Detect(buffer.Bytes())
	isVideo := strings.HasPrefix(strings.ToLower(mime.String()), "video/")
	fileLocation, err := aws.Upload(buffer, filePath, &class.Organization)
	if err != nil {
		awsDiagnostics.WriteAWSPartErr(err, "Upload Video Part")
		return "", false, err
	}
	return fileLocation, isVideo, nil
}

func DeleteVideoPart(part *videoParts.VideoPart, video *videoParts.Video, class *classesModel.Class) error {
	filePath := fmt.Sprintf("%s/videos/%s/%s", video.ClassHash, video.Hash, part.Name)
	return aws.Delete(filePath, class.Organization.S3BucketName)
}

func DeleteVideo(video *videoParts.Video, parts []videoParts.VideoPart, class *classesModel.Class) error {
	var err error
	for _, part := range parts {
		err = DeleteVideoPart(&part, video, class)
		if err != nil {
			return err
		}
	}
	folderPath := fmt.Sprintf("%s/videos/%s", video.ClassHash, video.Hash)
	return aws.Delete(folderPath, class.Organization.S3BucketName)
}
