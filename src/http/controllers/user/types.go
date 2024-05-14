package v1usercontroller

import (
	"database/sql"
	"net/http"
)

type V1User struct {
	DB *sql.DB
}

type iV1User interface {
	RegisterIt(http.ResponseWriter, *http.Request)
}

func New(v1User *V1User) iV1User {
	return v1User
}
