package v1usercontroller

import (
	"fmt"
	"net/http"

	"github.com/dwibi/health-record/src/helpers"
	userrepository "github.com/dwibi/health-record/src/repository/user"
	userusecase "github.com/dwibi/health-record/src/usecase/user"
	"github.com/go-playground/validator/v10"
)

type registerRequest struct {
	NIP      string `json:"nip" validate:"required,len=13"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (u V1User) RegisterIt(w http.ResponseWriter, r *http.Request) {
	payload := new(registerRequest)

	// check payload
	if err := helpers.ParseJSON(r, &payload); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	// check validation
	if err := helpers.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: fmt.Errorf("invalid payload: %v", errors).Error()})
		return
	}

	// create usecase
	uu := userusecase.New(
		userrepository.New(u.DB),
	)

	// use usecase to register IT user
	data, err := uu.RegisterIt(&userusecase.ParamsRegisterUserIt{
		NIP:      payload.NIP,
		Name:     payload.Name,
		Password: payload.Password,
	})

	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	helpers.WriteJSON(w, http.StatusCreated, SuccessResponse{
		Message: "User registered successfully",
		Data:    data,
	})
}
