package aws

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"net/http"
	"os"
	"strings"
)

// to load the region automatically from the environment variable
// the following line was added at the end of the ~/.profile file
// export AWS_REGION=eu-central-1
// while this line was added to the script at the server (Environment="AWS_REGION=eu-central-1")
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
	split := strings.Split(output.Location, ".com")
	cloudfrontLocation := "https://" + os.Getenv("CLOUDFRONT_DOMAIN") + split[1]
	cloudfrontLocation = strings.ReplaceAll(cloudfrontLocation, " ", "%20")
	cloudfrontLocation = strings.ReplaceAll(cloudfrontLocation, "+", "-")
	cloudfrontLocation = strings.ReplaceAll(cloudfrontLocation, "=", "_")
	return cloudfrontLocation, nil
}

