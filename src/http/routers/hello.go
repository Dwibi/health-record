package routes

import (
	v1helloController "github.com/dwibi/health-record/src/http/controllers/hello"
)

func (r *RouterTest) RegisterHello() {
	helloController := v1helloController.New(&v1helloController.V1hello{
		DB: r.DB,
	})

	r.Router.HandleFunc("/hello", helloController.FindOne).Methods("GET")
}
