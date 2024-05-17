package entities

import "time"

type MedicalRecords struct {
	ID             int64     `db:"id" json:"id"`
	IdentityNumber string    `db:"identity_number" json:"identity_number"`
	Symptoms       string    `db:"symptoms" json:"symptoms"`
	Medication     string    `db:"medication" json:"medication"`
	CreatedBy      string    `db:"created_by" json:"createdBy"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

type RecordSearchFilter struct {
	IdentityNumber int
	UserId         string
	Nip            string
	Limit          int
	Offset         int
	CreatedAt      string
}
