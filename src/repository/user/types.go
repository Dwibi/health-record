package userrepository

import (
	"database/sql"

	"github.com/dwibi/health-record/src/entities"
)

type sUserRepository struct {
	DB *sql.DB
}

type IUserRepository interface {
	CreateUserIt(*ParamsCreateUserIt) (*entities.User, error)
	IsExist(string) (bool, error)
	// CreateUserIt(*ParamsCreateUser) (*entities.User, error)
	// FindOne(*entities.ParamsCreateUser) (*entities.User, error)
	// IsExists(*entities.ParamsCreateUser) (bool, error)
}

func New(db *sql.DB) IUserRepository {
	return &sUserRepository{DB: db}
}
