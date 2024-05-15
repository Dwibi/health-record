package v1usercontroller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dwibi/health-record/src/helpers"
	userrepository "github.com/dwibi/health-record/src/repository/user"
	userusecase "github.com/dwibi/health-record/src/usecase/user"
	"github.com/go-playground/validator/v10"
)

func (u V1User) LoginNurse(w http.ResponseWriter, r *http.Request) {
	payload := new(loginRequest)

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

	nipStr := strconv.Itoa(payload.NIP)

	// TODO: Validate NIP
	if err := helpers.ValidateNIP(nipStr, "nurse"); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	// create usecase
	uu := userusecase.New(
		userrepository.New(u.DB),
	)

	// use usecase to login IT user
	data, status, err := uu.LoginNurse(&userusecase.ParamsLoginUserIt{
		NIP:      nipStr,
		Password: payload.Password,
	})

	if err != nil {
		helpers.WriteJSON(w, status, ErrorResponse{Message: err.Error()})
		return
	}

	helpers.WriteJSON(w, status, SuccessResponse{
		Message: "User login successfully",
		Data:    data,
	})
}
