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

type loginRequest struct {
	NIP      int    `json:"nip" validate:"required,number"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}

func (u V1User) LoginIt(w http.ResponseWriter, r *http.Request) {
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
	if err := helpers.ValidateNIP(nipStr); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	if isIT := helpers.IsItUser(nipStr); !isIT {
		helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: "nip should start 615"})
		return
	}
	// create usecase
	uu := userusecase.New(
		userrepository.New(u.DB),
	)

	// use usecase to login IT user
	data, status, err := uu.LoginIt(&userusecase.ParamsLoginUserIt{
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
