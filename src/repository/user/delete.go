package userrepository

import (
	"time"
)

func (i *sUserRepository) Delete(userId int) error {
	_, err := i.DB.Exec("UPDATE users SET updated_at = $1, deleted_at = $2 WHERE id = $3;", time.Now(), time.Now(), userId)

	if err != nil {
		return err
	}

	return nil
}
