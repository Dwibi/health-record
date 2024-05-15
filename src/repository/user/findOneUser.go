package userrepository

import (
	"database/sql"

	"github.com/dwibi/health-record/src/entities"
)

type ParamsFindOneUser struct {
	NIP string
}

func (i *sUserRepository) FindOneUser(p *ParamsFindOneUser) (*entities.User, error) {
	u := &entities.User{}
	err := i.DB.QueryRow("SELECT id, nip, name, password FROM users WHERE nip = $1 LIMIT 1", p.NIP).Scan(&u.ID, &u.NIP, &u.Name, &u.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return u, nil
}
