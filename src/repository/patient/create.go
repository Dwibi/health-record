package patientrepository

type ParamsCreate struct {
	IdentityNumber      int
	PhoneNumber         string
	Name                string
	Gender              string
	IdentityCardScanImg string
}

func (i *sPatientRepository) Create(p *ParamsCreate) error {
	/*
		INSERT INTO medical_patients (identity_number, phone_number, name, birth_date, gender, identity_card_scan_img
		) VALUES (1234567898765432, '+6285156055875', 'John Doe', '1990-01-15','male', 'https://http.cat/images/402.jpg');
	*/

	_, err := i.DB.Exec("INSERT INTO medical_patient (identity_number, phone_number, name, gender, identity_card_scan_img) VALUES ($1, $2, $3, $4, $5)", p.IdentityNumber, p.PhoneNumber, p.Name, p.Gender, p.IdentityCardScanImg)

	if err != nil {
		return err
	}

	return nil
}
