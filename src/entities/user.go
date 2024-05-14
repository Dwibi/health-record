package entities

import "time"

type User struct {
	ID                  int64     `json:"id"`
	NIP                 string    `json:"nip"`
	Name                string    `json:"name"`
	Password            string    `json:"password"`
	IdentityCardScanImg string    `json:"identityCardScanImg"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	DeletedAt           time.Time `json:"deleted_at"`
}
