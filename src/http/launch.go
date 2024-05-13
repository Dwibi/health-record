package http

import (
	"database/sql"
	"log"
	"net/http"

	routes "github.com/dwibi/health-record/src/http/routers"
	"github.com/gorilla/mux"
)

type Http struct {
	DB *sql.DB
}

type iHttp interface {
	Launch() error
}

func New(Http *Http) iHttp {
	return Http
}

func (h *Http) Launch() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/v1").Subrouter()

	v1 := routes.New(&routes.RouterTest{
		Router: subrouter,
		DB:     h.DB,
	})

	v1.RegisterHello()

	log.Println("Listening on", ":8080")

	return http.ListenAndServe(":8080", router)
}
