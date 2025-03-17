package s3

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var (
	s3Client   *s3.Client
	bucketName string
)

// Setup S3 Client
func SetupClient(region, bucket string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return fmt.Errorf("failed to load AWS config: %w", err)
	}

	s3Client = s3.NewFromConfig(cfg)
	bucketName = bucket
	return nil
}

// Setup S3 Client
func SetupClientWithStaticCreds(region, bucket, accessKeyID, secretAccessKey, endpoint string) error {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(aws.NewCredentialsCache(
			credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, ""),
		)),
		config.WithRegion(region),
		config.WithBaseEndpoint(endpoint),
	)
	if err != nil {
		return fmt.Errorf("failed to load AWS config: %w", err)
	}

	s3Client = s3.NewFromConfig(cfg)
	bucketName = bucket
	return nil
}

// UploadFile uploads a file to S3 from a multipart form.
func UploadFile(ctx context.Context, file *multipart.FileHeader, path string, acl types.ObjectCannedACL) (string, error) {
	// Open file content
	fileContent, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer fileContent.Close()

	// Create the object key (path + filename)
	objectKey := path + file.Filename

	// Create put object input
	putObjectInput := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   fileContent,
		ACL:    acl,
	}

	// Initialize the uploader
	uploader := manager.NewUploader(s3Client)

	// Upload the file
	result, err := uploader.Upload(ctx, putObjectInput)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %w", err)
	}

	// Return the file URL
	return result.Location, nil
}
