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
	LoginIt(http.ResponseWriter, *http.Request)
	LoginNurse(http.ResponseWriter, *http.Request)
	RegisterNurse(http.ResponseWriter, *http.Request)
	UpdateNurse(http.ResponseWriter, *http.Request)
	UpdateNurseAccess(http.ResponseWriter, *http.Request)
	DeleteNurse(http.ResponseWriter, *http.Request)
	FindMany(w http.ResponseWriter, r *http.Request)
}

func New(v1User *V1User) iV1User {
	return v1User
}
