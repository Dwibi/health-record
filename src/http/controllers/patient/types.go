package v1patientcontroller

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

type V1Patient struct {
	DB *sql.DB
}

type iVPatient interface {
	Create(http.ResponseWriter, *http.Request)
	FindMany(http.ResponseWriter, *http.Request)
}

func New(v1Patient *V1Patient) iVPatient {
	return v1Patient
}
