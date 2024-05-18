package recordrepository

import (
	"database/sql"

	"github.com/dwibi/health-record/src/entities"
)

type sRecordRepository struct {
	DB *sql.DB
}

type IRecordRepository interface {
	Create(*ParamsCreate) error
	FindMany(filters *entities.RecordSearchFilter) ([]*ResultFindMany, error)
}

func New(db *sql.DB) IRecordRepository {
	return &sRecordRepository{DB: db}
}
