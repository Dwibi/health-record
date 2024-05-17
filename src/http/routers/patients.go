package routes

import (
	v1patientcontroller "github.com/dwibi/health-record/src/http/controllers/patient"
	"github.com/dwibi/health-record/src/http/middleware"
)

func (r *RouterTest) RegisterPatient() {
	patientController := v1patientcontroller.New(&v1patientcontroller.V1Patient{
		DB: r.DB,
	})

	r.Router.HandleFunc("/medical/patient", middleware.AuthMiddleware(patientController.Create)).Methods("POST")
	r.Router.HandleFunc("/medical/patient", middleware.AuthMiddleware(patientController.FindMany)).Methods("GET")
}
