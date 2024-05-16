package userrepository

import "time"

type ParamsUpdatePassword struct {
	Id       int
	Password string
}

func (i *sUserRepository) UpdatePassword(p *ParamsUpdatePassword) error {
	_, err := i.DB.Exec("UPDATE users SET password = $1, updated_at = $2 WHERE id = $3", p.Password, time.Now(), p.Id)

	if err != nil {
		return err
	}

	return nil
}
