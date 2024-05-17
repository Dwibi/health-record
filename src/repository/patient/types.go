package patientrepository

import (
	"database/sql"

	"github.com/dwibi/health-record/src/entities"
)

type sPatientRepository struct {
	DB *sql.DB
}

type IPatientRepository interface {
	Create(*ParamsCreate) error
	FindMany(*entities.PatientSearchFilter) ([]*ResultFindMany, error)
}

func New(db *sql.DB) IPatientRepository {
	return &sPatientRepository{DB: db}
}
