package patientrepository

import "log"

func (i *sPatientRepository) IsExist(identityNumber string) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM medical_patients WHERE identity_number = $1);"
	var exists bool

	err := i.DB.QueryRow(query, identityNumber).Scan(&exists)

	if err != nil {
		log.Printf("Error checking if medical_patients exists %v", err)
		return false, err
	}
	return exists, nil
}
