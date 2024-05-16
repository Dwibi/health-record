package userrepository

import (
	"log"

	"github.com/dwibi/health-record/src/entities"
)

type ParamsCreateUser struct {
	NIP                 string
	Name                string
	Password            string
	IdentityCardScanImg string
}

func (i *sUserRepository) CreateUser(p *ParamsCreateUser) (*entities.User, error) {
	var id int
	err := i.DB.QueryRow("INSERT INTO users (nip, name, password, identity_card_scan_img) VALUES ($1, $2, $3, $4) RETURNING id;", p.NIP, p.Name, p.Password, p.IdentityCardScanImg).Scan(&id)

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
