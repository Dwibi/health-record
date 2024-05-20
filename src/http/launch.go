package http

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	routes "github.com/dwibi/health-record/src/http/routers"
	"github.com/gorilla/mux"
)

type Http struct {
	DB       *sql.DB
	Uploader *manager.Uploader
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

	// subrouter.Use(middleware.LoggingMiddleware)

	v1 := routes.New(&routes.RouterTest{
		Router:   subrouter,
		DB:       h.DB,
		Uploader: h.Uploader,
	})

	v1.RegisterHello()
	v1.RegisterUser()
	v1.RegisterPatient()
	v1.RegisterRecord()
	v1.RegisterUpload()

	log.Println("Listening on", ":8080")

	return http.ListenAndServe(":8080", router)
}
