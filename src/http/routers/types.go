package routes

import (
	"database/sql"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/gorilla/mux"
)

type RouterTest struct {
	Router   *mux.Router
	DB       *sql.DB
	Uploader *manager.Uploader
}

type iRoutes interface {
	RegisterHello()
	RegisterUser()
	RegisterPatient()
	RegisterRecord()
	RegisterUpload()
}

func New(routes *RouterTest) iRoutes {
	return routes
}
