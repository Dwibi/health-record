package recordrepository

import (
	"database/sql"
)

type sRecordRepository struct {
	DB *sql.DB
}

type IRecordRepository interface {
	Create(*ParamsCreate) error
	// FindMany(*entities.PatientSearchFilter) ([]*ResultFindMany, error)
}

func New(db *sql.DB) IRecordRepository {
	return &sRecordRepository{DB: db}
}
