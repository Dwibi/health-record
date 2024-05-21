package v1uploadcontroller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/dwibi/health-record/src/helpers"
	userrepository "github.com/dwibi/health-record/src/repository/user"
	uploadusecase "github.com/dwibi/health-record/src/usecase/upload"
	"github.com/google/uuid"
)

type returnData struct {
	ImageUrl string `json:"imageUrl"`
}

func (u V1Upload) Image(w http.ResponseWriter, r *http.Request) {
	// Access the user information from the request context
	userIdClaims, ok := r.Context().Value(helpers.UserContextKey).(int)
	if !ok {
		http.Error(w, "User information not found in context", http.StatusInternalServerError)
		return
	}

	// Set a maximum upload size
	r.ParseMultipartForm(2 << 20)

	// Get the file from the request
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate the file size
	if handler.Size < 10*1024 || handler.Size > 2*1024*1024 {
		http.Error(w, "File size must be between 10KB and 2MB", http.StatusBadRequest)
		return
	}

	// Validate the file type
	fileType := handler.Header.Get("Content-Type")
	if fileType != "image/jpeg" && fileType != "image/jpg" {
		http.Error(w, "File must be in *.jpg or *.jpeg format", http.StatusBadRequest)
		return
	}

	uu := uploadusecase.New(
		userrepository.New(u.DB),
	)

	status, err := uu.ValidateRole(userIdClaims)

	if err != nil {
		helpers.WriteJSON(w, status, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	// Generate a UUID for the file name
	uuid := uuid.New().String()
	ext := filepath.Ext(handler.Filename)
	newFileName := fmt.Sprintf("%s%s", uuid, ext)

	// Define S3 bucket and key
	bucket := os.Getenv("AWS_S3_BUCKET_NAME")
	key := newFileName

	result, err := u.Uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		ACL:    "public-read",
		Body:   file,
	})

	if err != nil {
		http.Error(w, "Error S3", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, SuccessResponse{
		Message: "File uploaded sucessfully",
		Data: returnData{
			ImageUrl: result.Location,
		},
	})
}
