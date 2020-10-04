package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)
// to load the region automatically from the config file in ~/.aws/config
// the following line was added at the end of the ~/.profile file
// export AWS_SDK_LOAD_CONFIG=true
func ListBuckets() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println(err)
	}
	s3service := s3.New(sess)
	result, err := s3service.ListBuckets(nil)
	if err != nil {
		fmt.Println(err)
	}
	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}
