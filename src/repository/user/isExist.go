package userrepository

import "log"

func (i *sUserRepository) IsExist(nip string) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE nip = $1);"
	var exists bool

	err := i.DB.QueryRow(query, nip).Scan(&exists)

	if err != nil {
		log.Printf("Error checking if user exists %v", err)
		return false, err
	}
	return exists, nil
}
