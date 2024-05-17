package entities

import (
	"time"
)

type GenderEnum string

const (
	Male   GenderEnum = "male"
	Female GenderEnum = "female"
)

type MedicalPatient struct {
	ID                  int64     `db:"id" json:"id"`
	IdentityNumber      string    `db:"identity_number" json:"identity_number"`
	PhoneNumber         string    `db:"phone_number" json:"phone_number"`
	Name                string    `db:"name" json:"name"`
	BirthDate           time.Time `db:"birth_date" json:"birth_date"`
	Gender              string    `db:"gender" json:"gender"`
	IdentityCardScanImg string    `db:"identity_card_scan_img" json:"identity_card_scan_img"`
	CreatedAt           time.Time `db:"created_at" json:"created_at"`
	UpdatedAt           time.Time `db:"updated_at" json:"updated_at"`
}

type PatientSearchFilter struct {
	IdentityNumber int    //
	Name           string //
	PhoneNumber    int    //
	Limit          int    //
	Offset         int    //
	CreatedAt      string //
}
