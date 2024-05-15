package routes

import (
	v1usercontroller "github.com/dwibi/health-record/src/http/controllers/user"
	"github.com/dwibi/health-record/src/http/middleware"
)

func (r *RouterTest) RegisterUser() {
	usersController := v1usercontroller.New(&v1usercontroller.V1User{
		DB: r.DB,
	})

	r.Router.HandleFunc("/user/it/register", usersController.RegisterIt).Methods("POST")
	r.Router.HandleFunc("/user/it/login", usersController.LoginIt).Methods("POST")
	r.Router.HandleFunc("/user/nurse/login", usersController.LoginNurse).Methods("POST")
	r.Router.HandleFunc("/user/nurse/register", middleware.AuthMiddleware(usersController.RegisterNurse)).Methods("POST")

}
