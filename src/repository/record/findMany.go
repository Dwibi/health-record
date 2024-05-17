package recordrepository

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/dwibi/health-record/src/entities"
)

type ResultFindMany struct {
	IdentityNumber int       `json:"identityNumber"`
	PhoneNumber    string    `json:"phoneNumber"`
	Name           string    `json:"name"`
	BirthDate      time.Time `json:"birthDate"`
	Gender         string    `json:"gender"`
	CreatedAt      time.Time `json:"createdAt"`
}

func (i *sPatientRepository) FindMany(filters *entities.PatientSearchFilter) ([]*ResultFindMany, error) {
	query := "SELECT identity_number, phone_number, name, birth_date, gender, created_at FROM users WHERE 1=1"
	params := []interface{}{}

	n := (&entities.PatientSearchFilter{})

	if !reflect.DeepEqual(filters, n) {
		conditions := []string{}

		if filters.IdentityNumber != 0 {
			conditions = append(conditions, "identity_number = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.IdentityNumber)
		}

		if filters.PhoneNumber != 0 {
			conditions = append(conditions, "id LIKE $"+strconv.Itoa(len(params)+1))
			params = append(params, "%"+strconv.Itoa(filters.PhoneNumber)+"%")
		}

		if filters.Name != "" {
			conditions = append(conditions, "lower(name) LIKE lower($"+strconv.Itoa(len(params)+1)+")")
			params = append(params, "%"+filters.Name+"%")
		}

		if len(conditions) > 0 {
			query += " AND "
		}
		query += strings.Join(conditions, " AND ")
	}

	if filters.CreatedAt != "" {
		if filters.CreatedAt == "desc" {
			query += " ORDER BY created_at DESC"
		}
		if filters.CreatedAt == "asc" {
			query += " ORDER BY created_at ASC"
		}
	} else {
		query += " ORDER BY created_at ASC"
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

	fmt.Println(query)

	rows, err := i.DB.Query(query, params...)
	if err != nil {
		log.Printf("Error finding cat: %s", err)
		return nil, err
	}
	defer rows.Close()

	users := make([]*ResultFindMany, 0)
	for rows.Next() {
		c := new(ResultFindMany)
		var identityNum string
		err := rows.Scan(&identityNum, &c.PhoneNumber, &c.Name, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		c.IdentityNumber = func() int { n, _ := strconv.Atoi(identityNum); return n }()
		fmt.Println(c)
		users = append(users, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
