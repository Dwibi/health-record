package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type RouterTest struct {
	Router *mux.Router
	DB     *sql.DB
}

type iRoutes interface {
	RegisterHello()
}

func New(routes *RouterTest) iRoutes {
	return routes
}
