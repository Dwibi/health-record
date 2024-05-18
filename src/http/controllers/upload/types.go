package v1uploadcontroller

import (
	"database/sql"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type V1Upload struct {
	DB       *sql.DB
	Uploader *manager.Uploader
}

type iVupload interface {
	Image(http.ResponseWriter, *http.Request)
	// Create(http.ResponseWriter, *http.Request)
	// FindMany(http.ResponseWriter, *http.Request)
}

func New(v1Upload *V1Upload) iVupload {
	return v1Upload
}
