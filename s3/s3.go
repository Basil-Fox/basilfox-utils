package s3

import (
	"context"
	"mime/multipart"

	"github.com/FiberApps/common-library/logger"
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
func SetupClient(region, bucket string) {
	log := logger.New()

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Error("AWS_S3:: Error while loading config: %v", err)
	}

	s3Client = s3.NewFromConfig(cfg)
	bucketName = bucket
}

// Upload file from multipart form
func UploadFile(file *multipart.FileHeader) (string, error) {
	log := logger.New()

	fileReader, err := file.Open()
	if err != nil {
		log.Error("AWS_S3:: Error while opening file: %v", err)
		return "", err
	}

	uploader := manager.NewUploader(s3Client)

	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(file.Filename),
		Body:   fileReader,
		ACL:    "public-read",
	})
	if err != nil {
		log.Error("AWS_S3:: Error while uploading file: %v", err)
		return "", err
	}

	// Generate the file URL
	return result.Location, nil
}
