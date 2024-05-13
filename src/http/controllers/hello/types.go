package v1helloController

import (
	"database/sql"
	"net/http"
)

type V1hello struct {
	DB *sql.DB
}

type iV1Hello interface {
	FindOne(http.ResponseWriter, *http.Request)
}

func New(v1Hello *V1hello) iV1Hello {
	return v1Hello
}
