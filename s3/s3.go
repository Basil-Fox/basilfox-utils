package s3

import (
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	s3Client   *s3.Client
	bucketName string
)

// Setup S3 Client
func SetupClient(region, bucket string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return err
	}

	s3Client = s3.NewFromConfig(cfg)
	bucketName = bucket
	return nil
}

// Upload file from multipart form
func UploadFile(file *multipart.FileHeader, path string) (string, error) {
	fileContent, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileContent.Close()

	// Create put object input
	putObjectInput := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path + file.Filename),
		Body:   fileContent,
		ACL:    "public-read",
	}

	uploader := manager.NewUploader(s3Client)

	result, err := uploader.Upload(context.TODO(), putObjectInput)
	if err != nil {
		return "", err
	}

	// Generate the file URL
	return result.Location, nil
}
