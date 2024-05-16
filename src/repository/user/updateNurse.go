package userrepository

import "time"

type ParamsUpdateNurse struct {
	Id   int
	NIP  string
	Name string
}

func (i *sUserRepository) UpdateNurse(p *ParamsUpdateNurse) error {
	_, err := i.DB.Exec("UPDATE users SET NIP = $1, name = $2, updated_at = $3 WHERE id = $4", p.NIP, p.Name, time.Now(), p.Id)

	if err != nil {
		return err
	}

	return nil
}
