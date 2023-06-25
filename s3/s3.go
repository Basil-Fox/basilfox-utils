package s3

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadFile(fileBuffer *bytes.Buffer, fileName string, fileType string) (string, error) {
	var (
		accessKeyID     = os.Getenv("S3_ACCESS_KEY")
		secretAccessKey = os.Getenv("S3_SECRET_KEY")
		region          = os.Getenv("S3_REGION")
		bucketName      = os.Getenv("S3_BUCKET")
	)

	// Create a new session with default session credentials
	s3session, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		return "", err
	}

	// Create a new S3 service client
	svc := s3.New(s3session)

	// Configure the S3 object input parameters
	input := &s3.PutObjectInput{
		Bucket:        aws.String(bucketName),
		Key:           aws.String(fileName),
		Body:          bytes.NewReader(fileBuffer.Bytes()),
		ContentLength: aws.Int64(int64(fileBuffer.Len())),
		ContentType:   aws.String(fileType),
	}

	// Upload the image to S3
	_, err = svc.PutObject(input)
	if err != nil {
		return "", err
	}

	// Generate the file URL
	return fmt.Sprintf("https://%s.s3-%s.amazonaws.com/%s", bucketName, region, fileName), nil

}
