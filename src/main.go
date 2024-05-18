package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dwibi/health-record/src/http"
	"github.com/dwibi/health-record/src/http/drivers/db"
	"github.com/joho/godotenv"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	godotenv.Load()
	dbConnection, err := db.CreateConnection()

	if err != nil {
		fmt.Println("Error creating database connection:", err)
		return
	}

	defer func() {
		if err := dbConnection.Close(); err != nil {
			fmt.Println("Error closing database connection:", err)
		}
	}()

	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)

	h := http.New(&http.Http{
		DB:       dbConnection,
		Uploader: uploader,
	})

	h.Launch()
}
