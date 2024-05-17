package recordrepository

type ParamsCreate struct {
	CreatedBy      int
	IdentityNumber string
	Symptoms       string
	Medications    string
}

func (i *sRecordRepository) Create(p *ParamsCreate) error {
	_, err := i.DB.Exec("INSERT INTO medical_records (iidentity_number, symptoms, medications, created_by) VALUES ($1, $2, $3, $4)", p.IdentityNumber, p.Symptoms, p.Medications, p.CreatedBy)

	if err != nil {
		return err
	}

	return nil
}
