Here is a README file template for your **BasilFox Common Library** that you can use for your microservices:

---

# BasilFox Common Library

This repository contains the common utility library for microservices used across the **BasilFox** ecosystem. It includes reusable functions and modules to streamline tasks such as Firebase integration, Redis, S3 file uploads, database connections, push notifications, and more. The library is built using Go and designed to work seamlessly with the Fiber web framework.

## Features

- **Firebase Integration**: Simplified methods for interacting with Firebase services such as authentication and push notifications.
- **Redis Client**: A wrapper for Redis to easily connect and perform operations with the database.
- **S3 File Upload**: Functions to upload files to AWS S3 with minimal configuration.
- **Database (MongoDB)**: Connection management for MongoDB with easy ping and connection verification.
- **Push Notifications**: Supports both regular and silent notifications for multiple tokens using Firebase Cloud Messaging (FCM).
- **Middleware**: Includes utility middleware for request handling like request ID generation and header validation.
- **Environment Management**: Loads `.env` configurations for non-production environments.

## Installation

To use the **BasilFox Common Library** in your own microservices, simply add it to your Go project as a module dependency.

```bash
go get github.com/Basil-Fox/basilfox-utils
```

Make sure you have Go modules enabled in your project by running `go mod init <module-name>` if it’s not already set up.

## Usage

### Firebase Setup

1. Initialize Firebase with a credentials file or JSON:

```go
package firebase

import (
	"log"
	"github.com/Basil-Fox/basilfox-utils/firebase"
)

func initFirebase() {
	err := firebase.InitWithFile("path/to/credentials.json")
	if err != nil {
		log.Fatalf("Firebase initialization failed: %v", err)
	}
}
```

2. Send push notifications to a list of tokens:

```go
package firebase

import (
	"log"
	"github.com/Basil-Fox/basilfox-utils/firebase"
)

func sendNotification() {
	msg := kafka.SendPushNotificationMessage{
		Title:   "Test Title",
		Body:    "Test Body",
		Tokens:  []string{"token1", "token2"},
		Data:    map[string]string{"key": "value"},
	}

	err := firebase.SendToTokens(msg, false)
	if err != nil {
		log.Printf("Error sending push notifications: %v", err)
	}
}
```

### Redis Setup

1. Connect to Redis:

```go
package redis

import (
	"log"
	"github.com/Basil-Fox/basilfox-utils/redis"
)

func initRedis() {
	err := redis.Connect("localhost:6379", "user", "password")
	if err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}
}
```

### S3 File Upload

1. Set up the S3 client and upload a file:

```go
package s3

import (
	"log"
	"mime/multipart"
	"github.com/Basil-Fox/basilfox-utils/s3"
)

func uploadFile(file *multipart.FileHeader) {
	err := s3.SetupClient("us-east-1", "your-bucket-name")
	if err != nil {
		log.Fatalf("S3 client setup failed: %v", err)
	}

	url, err := s3.UploadFile(file, "uploads/")
	if err != nil {
		log.Printf("Error uploading file: %v", err)
	} else {
		log.Printf("File uploaded successfully. URL: %s", url)
	}
}
```

### Database (MongoDB) Connection

1. Connect to MongoDB:

```go
package database

import (
	"log"
	"github.com/Basil-Fox/basilfox-utils/database"
)

func initDatabase() {
	err := database.Connect("mongodb://localhost:27017", "your-database-name")
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}
}
```

### Middleware

1. Use middleware to validate headers in your Fiber app:

```go
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Basil-Fox/basilfox-utils/middleware"
)

func main() {
	app := fiber.New()

	// Apply middleware
	app.Use(middleware.ValidateHeaders(middleware.EndpointPrivate, []string{"namespace1", "namespace2"}))

	app.Listen(":3000")
}
```

### Environment Variables

The library uses `godotenv` to load environment variables. Make sure to create a `.env` file for non-production environments.

```env
# .env
DATABASE_URL=mongodb://localhost:27017
REDIS_URI=localhost:6379
S3_BUCKET_NAME=your-bucket-name
```

## Contributing

We welcome contributions! If you have an improvement or bug fix, feel free to open a pull request. Please ensure that your code is well-tested and follows the Go conventions.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

### Notes:

- Replace the specific usage examples with the relevant code from your library.
- Add more sections as needed, such as troubleshooting, testing, or other configurations that might be helpful to your users.
