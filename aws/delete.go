package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Delete(filePath, s3BucketName string) error {
	sess, err := session.NewSession()
	if err != nil {
		return err
	}
	// Create S3 service client
	svc := s3.New(sess)
	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s3BucketName),
		Key: aws.String(filePath),
	})
	if err != nil {
		return err
	}
	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(s3BucketName),
		Key: aws.String(filePath),
	})
	return err
}

func DeleteFolder(folderPath, s3BucketName string) error {
	sess, err := session.NewSession()
	if err != nil {
		return err
	}
	svc := s3.New(sess)
	iter := s3manager.NewDeleteListIterator(svc, &s3.ListObjectsInput{
		Bucket: aws.String(s3BucketName),
		Prefix: aws.String(folderPath),
	})
	return s3manager.NewBatchDeleteWithClient(svc).Delete(aws.BackgroundContext(), iter)
}