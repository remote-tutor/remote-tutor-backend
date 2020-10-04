package aws

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"net/http"
	"os"
)

// to load the region automatically from the config file in ~/.aws/config
// the following line was added at the end of the ~/.profile file
// export AWS_SDK_LOAD_CONFIG=true
func Upload(buffer *bytes.Buffer, filePath string) (string, error) {
	contentType := http.DetectContentType(buffer.Bytes())
	sess, err := session.NewSession()
	if err != nil {
		return "", err
	}
	uploader := s3manager.NewUploader(sess)
	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("BUCKET_NAME")),
		Key:         aws.String(filePath),
		Body:        buffer,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", err
	}
	return output.Location, nil
}

