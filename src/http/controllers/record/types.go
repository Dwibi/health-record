package v1recordcontroller

import (
	"database/sql"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type V1Record struct {
	DB *sql.DB
}

type iVRecord interface {
	Create(http.ResponseWriter, *http.Request)
	// FindMany(http.ResponseWriter, *http.Request)
}

func New(v1Record *V1Record) iVRecord {
	return v1Record
}
