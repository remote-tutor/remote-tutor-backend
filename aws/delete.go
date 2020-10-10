package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func Delete(filePath string) error {
	sess, err := session.NewSession()
	if err != nil {
		return err
	}
	// Create S3 service client
	svc := s3.New(sess)
	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key: aws.String(filePath),
	})
	if err != nil {
		return err
	}
	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key: aws.String(filePath),
	})
	return err
}