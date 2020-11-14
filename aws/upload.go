package aws

import (
	organizationsModel "backend/models/organizations"
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gabriel-vasile/mimetype"
	"strings"
)

// to load the region automatically from the environment variable
// the following line was added at the end of the ~/.profile file
// export AWS_REGION=eu-central-1
func Upload(buffer *bytes.Buffer, filePath string, organization *organizationsModel.Organization) (string, error) {
	mime := mimetype.Detect(buffer.Bytes())
	sess, err := session.NewSession()
	if err != nil {
		return "", err
	}
	uploader := s3manager.NewUploader(sess)
	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(organization.S3BucketName),
		Key:         aws.String(filePath),
		Body:        buffer,
		ContentType: aws.String(mime.String()),
	})
	if err != nil {
		return "", err
	}
	split := strings.Split(output.Location, ".com")
	cloudfrontLocation := "https://" + organization.CloudfrontDomain + split[1]
	cloudfrontLocation = strings.ReplaceAll(cloudfrontLocation, " ", "%20")
	cloudfrontLocation = strings.ReplaceAll(cloudfrontLocation, "+", "-")
	cloudfrontLocation = strings.ReplaceAll(cloudfrontLocation, "=", "_")
	return cloudfrontLocation, nil
}

