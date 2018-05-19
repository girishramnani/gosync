package gosync

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3BucketUploader struct {
	Bucket   string
	session  *session.Session
	s3client *s3.S3
}

func NewS3BucketUploader(bucket string) *S3BucketUploader {

	session := makeSession("default")
	s3client := s3.New(session)

	return &S3BucketUploader{
		Bucket:   bucket,
		session:  session,
		s3client: s3client,
	}

}

func (s3b *S3BucketUploader) Upload(filePath string) error {

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	// Upload the file to the s3 given bucket
	params := &s3.PutObjectInput{
		Bucket: aws.String(s3b.Bucket), // Required
		Key:    aws.String(filePath),   // Required
		Body:   file,
	}
	_, err = s3b.s3client.PutObject(params)
	if err != nil {
		return err
	}
	return nil
}

func makeSession(profile string) *session.Session {
	// Enable loading shared config file
	os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
	// Specify profile to load for the session's config
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: profile,
	})
	if err != nil {
		fmt.Println("failed to create session,", err)
		fmt.Println(err)
		os.Exit(1)
	}

	return sess
}
