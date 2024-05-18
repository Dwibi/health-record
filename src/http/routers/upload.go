package routes

import (
	v1uploadcontroller "github.com/dwibi/health-record/src/http/controllers/upload"
	"github.com/dwibi/health-record/src/http/middleware"
)

func (r *RouterTest) RegisterUpload() {
	uploadController := v1uploadcontroller.New(&v1uploadcontroller.V1Upload{
		Uploader: r.Uploader,
		DB:       r.DB,
	})

	r.Router.HandleFunc("/image", middleware.AuthMiddleware(uploadController.Image)).Methods("POST")
	// r.Router.HandleFunc("/medical/patient", middleware.AuthMiddleware(patientController.FindMany)).Methods("GET")
}
