package s3

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/FiberApps/core/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var s3Client *s3.Client

// Setup S3 Client
func SetupClient() {
	region := os.Getenv("S3_REGION")
	log := logger.New()

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	s3Client = s3.NewFromConfig(cfg)
}

// Upload file from multipart form
func UploadFile(file *multipart.FileHeader) (string, error) {
	bucketName := os.Getenv("S3_BUCKET")
	fileReader, err := file.Open()
	if err != nil {
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
		return "", err
	}

	// Generate the file URL
	return result.Location, nil
}
