package routes

import (
	v1recordcontroller "github.com/dwibi/health-record/src/http/controllers/record"
	"github.com/dwibi/health-record/src/http/middleware"
)

func (r *RouterTest) RegisterRecord() {
	recordController := v1recordcontroller.New(&v1recordcontroller.V1Record{
		DB: r.DB,
	})

	r.Router.HandleFunc("/medical/record", middleware.AuthMiddleware(recordController.Create)).Methods("POST")
	r.Router.HandleFunc("/medical/record", middleware.AuthMiddleware(recordController.FindMany)).Methods("GET")
}
