package v1recordcontroller

import (
	"fmt"
	"net/http"

	"github.com/dwibi/health-record/src/helpers"
	patientrepository "github.com/dwibi/health-record/src/repository/patient"
	recordrepository "github.com/dwibi/health-record/src/repository/record"
	userrepository "github.com/dwibi/health-record/src/repository/user"
	recordusecase "github.com/dwibi/health-record/src/usecase/record"
	"github.com/go-playground/validator/v10"
)

type createRequest struct {
	IdentityNumber int    `json:"identityNumber" validate:"required,number"`
	Symptoms       string `json:"symptoms" validate:"required,min=1,max=2000"`
	Medications    string `json:"medications" validate:"required,min=1,max=2000"`
}

func (u V1Record) Create(w http.ResponseWriter, r *http.Request) {
	// Access the user information from the request context
	userIdClaims, ok := r.Context().Value(helpers.UserContextKey).(int)
	if !ok {
		http.Error(w, "User information not found in context", http.StatusInternalServerError)
		return
	}

	// Validate the payload
	payload := new(createRequest)

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

	if isIdentityNum := helpers.ValidateIdentityNum(payload.IdentityNumber); !isIdentityNum {
		helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: "identityNumber should be 16 digits"})
		return
	}

	//create usecase
	uu := recordusecase.New(
		recordrepository.New(u.DB),
		patientrepository.New(u.DB),
		userrepository.New(u.DB),
	)

	// use usecase to create patient
	status, err := uu.Create(&recordusecase.ParamsCreate{
		ReqUserId:      userIdClaims,
		IdentityNumber: payload.IdentityNumber,
		Symptoms:       payload.Symptoms,
		Medications:    payload.Medications,
	})

	if err != nil {
		helpers.WriteJSON(w, status, ErrorResponse{Message: err.Error()})
		return
	}

	helpers.WriteJSON(w, status, SuccessResponse{
		Message: "Created!",
		Data:    nil,
	})
}
