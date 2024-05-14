package userrepository

import (
	"log"

	"github.com/dwibi/health-record/src/entities"
)

type ParamsCreateUserIt struct {
	NIP      string
	Name     string
	Password string
}

func (i *sUserRepository) CreateUserIt(p *ParamsCreateUserIt) (*entities.User, error) {
	var id int64
	err := i.DB.QueryRow("INSERT INTO users (nip, name, password) VALUES ($1, $2, $3) RETURNING id;", p.NIP, p.Name, p.Password).Scan(&id)

	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return nil, err
	}

	user := &entities.User{
		ID:   id,
		NIP:  p.NIP,
		Name: p.Name,
	}

	return user, nil
}
