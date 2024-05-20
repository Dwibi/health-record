package recordrepository

import (
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/dwibi/health-record/src/entities"
)

type identityDetail struct {
	IdentityNumber      int       `json:"identityNumber"`
	PhoneNumber         string    `json:"phoneNumber"`
	Name                string    `json:"name"`
	BirthDate           time.Time `json:"birthDate"`
	Gender              string    `json:"gender"`
	IdentityCardScanImg string    `json:"identityCardScanImg"`
}

type createdBy struct {
	Nip    int    `json:"nip"`
	Name   string `json:"name"`
	UserId string `json:"userId"`
}
type ResultFindMany struct {
	IdentityDetail identityDetail `json:"identityDetail"`
	Symptoms       string         `json:"symptoms"`
	Medications    string         `json:"medications"`
	CreatedBy      createdBy      `json:"createdBy"`
	CreatedAt      time.Time      `json:"createdAt"`
}

func (i *sRecordRepository) FindMany(filters *entities.RecordSearchFilter) ([]*ResultFindMany, error) {
	query := "SELECT mp.identity_number AS identityNumber, mp.phone_number AS phoneNumber, mp.name AS patientName, mp.birth_date AS birthDate, mp.gender AS gender, mp.identity_card_scan_img AS identityCardScanImg, mr.symptoms AS symptoms, mr.medications AS medications, mr.created_at AS createdAt, u.nip AS nip, u.name AS createdByName, u.id AS userId FROM medical_records mr JOIN medical_patients mp ON mr.identity_number_patient = mp.identity_number JOIN users u ON mr.created_by = u.id WHERE 1=1"
	params := []interface{}{}

	n := (&entities.PatientSearchFilter{})

	if !reflect.DeepEqual(filters, n) {
		conditions := []string{}

		if filters.IdentityNumber != 0 {
			conditions = append(conditions, "mp.identity_number = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.IdentityNumber)
		}

		if filters.UserId != "" {
			conditions = append(conditions, "u.id = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.UserId)
		}

		if filters.Nip != "" {
			conditions = append(conditions, "u.nip LIKE $"+strconv.Itoa(len(params)+1))
			params = append(params, "%"+filters.Nip+"%")
		}

		if len(conditions) > 0 {
			query += " AND "
		}
		query += strings.Join(conditions, " AND ")
	}

	if filters.CreatedAt != "" {
		if filters.CreatedAt == "desc" {
			query += " ORDER BY mr.created_at DESC"
		}
		if filters.CreatedAt == "asc" {
			query += " ORDER BY mr.created_at ASC"
		}
	} else {
		query += " ORDER BY mr.created_at DESC"
	}

	if filters.Limit == 0 {
		filters.Limit = 5
	}

	query += " LIMIT $" + strconv.Itoa(len(params)+1)
	params = append(params, filters.Limit)

	if filters.Offset == 0 {
		filters.Offset = 0
	} else {
		query += " OFFSET $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.Offset)
	}

	rows, err := i.DB.Query(query, params...)
	if err != nil {
		log.Printf("Error finding cat: %s", err)
		return nil, err
	}
	defer rows.Close()

	data := make([]*ResultFindMany, 0)
	for rows.Next() {
		var (
			id        identityDetail
			cb        createdBy
			symptoms  string
			meds      string
			createdAt time.Time
		)
		err := rows.Scan(
			&id.IdentityNumber,
			&id.PhoneNumber,
			&id.Name,
			&id.BirthDate,
			&id.Gender,
			&id.IdentityCardScanImg,
			&symptoms,
			&meds,
			&createdAt,
			&cb.Nip,
			&cb.Name,
			&cb.UserId,
		)
		if err != nil {
			return nil, err
		}
		record := &ResultFindMany{
			IdentityDetail: id,
			Symptoms:       symptoms,
			Medications:    meds,
			CreatedBy:      cb,
			CreatedAt:      createdAt,
		}

		data = append(data, record)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
