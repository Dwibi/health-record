package userrepository

import (
	"database/sql"

	"github.com/dwibi/health-record/src/entities"
)

func (i *sUserRepository) GetUserById(id int) (*entities.User, error) {
	u := &entities.User{}
	err := i.DB.QueryRow("SELECT id, nip, name FROM users WHERE id = $1 AND deleted_at IS NULL", id).Scan(&u.ID, &u.NIP, &u.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return u, nil
}
