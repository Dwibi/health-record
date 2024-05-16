package userrepository

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
	UserId    string    `json:"userId"`
	Nip       int       `json:"nip"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	// "userId": "", // ID should be a string
	// "nip": 6152200102987
	// "name": "namadepan namabelakang",
	// "createdAt": "" // should in ISO 8601 format
}

func (i *sUserRepository) FindMany(filters *entities.UserSearchFilter) ([]*ResultFindMany, error) {
	query := "SELECT id, nip, name, created_at FROM users WHERE 1=1"
	params := []interface{}{}

	n := (&entities.UserSearchFilter{})

	if !reflect.DeepEqual(filters, n) {
		conditions := []string{}

		if filters.UserId != 0 {
			conditions = append(conditions, "id = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.UserId)
		}

		if filters.NIP != 0 {
			conditions = append(conditions, "id LIKE $"+strconv.Itoa(len(params)+1))
			params = append(params, "%"+strconv.Itoa(filters.NIP)+"%")
		}

		if filters.Name != "" {
			conditions = append(conditions, "lower(name) LIKE lower($"+strconv.Itoa(len(params)+1)+")")
			params = append(params, "%"+filters.Name+"%")
		}

		if filters.Role != "" {
			if filters.Role == "it" {
				conditions = append(conditions, "id LIKE $"+strconv.Itoa(len(params)+1))
				params = append(params, "615%")
			}
			if filters.Role == "nurse" {
				conditions = append(conditions, "id LIKE $"+strconv.Itoa(len(params)+1))
				params = append(params, "303%")
			}
		}

		if len(conditions) > 0 {
			query += " AND "
		}
		query += strings.Join(conditions, " AND ")
	}

	if filters.CreatedAt != "" {
		if filters.CreatedAt == "desc" {
			query += " ORDER BY created_at DESC"
		} else if filters.CreatedAt == "asc" {
			query += " ORDER BY created_at ASC"
		} else {
			query += " ORDER BY created_at ASC"

		}
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
		var userId int
		var nip string
		err := rows.Scan(&userId, &nip, &c.Name, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		c.UserId = strconv.Itoa(userId)
		c.Nip = func() int { n, _ := strconv.Atoi(nip); return n }()
		fmt.Println(c)
		users = append(users, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
