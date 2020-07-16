package aws

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Upload(fileName string, bucket string, key string, region string) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewSharedCredentials("/home/nirmal/.aws/credentials", "default"),
	}))
	s3session := s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	})))

	resp, err := s3session.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	uploader := s3manager.NewUploader(sess)
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
		return err
	}
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   f,
	})
	fmt.Println(result)
	if err != nil {
		panic(err)
		return err
	}
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	return nil
}
