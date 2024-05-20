package v1patientcontroller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dwibi/health-record/src/helpers"
	patientrepository "github.com/dwibi/health-record/src/repository/patient"
	userrepository "github.com/dwibi/health-record/src/repository/user"
	patientusecase "github.com/dwibi/health-record/src/usecase/patient"
	"github.com/go-playground/validator/v10"
)

type createRequest struct {
	IdentityNumber      int    `json:"identityNumber" validate:"required,number"`
	PhoneNumber         string `json:"phoneNumber" validate:"required,min=10,max=15"`
	Name                string `json:"name" validate:"required,min=3,max=30"`
	Gender              string `json:"gender" validate:"required"`
	BirthDate           string `json:"birthDate" validate:"required"`
	IdentityCardScanImg string `json:"identityCardScanImg" validate:"required,url"`
}

func (u V1Patient) Create(w http.ResponseWriter, r *http.Request) {
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

	if isPhoneNum := helpers.ValidateInaPhoneNum(payload.PhoneNumber); !isPhoneNum {
		helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: "phone number should start +62"})
		return
	}

	if isGenderValid := helpers.ValidateGender(payload.Gender); !isGenderValid {
		helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: "should be 'male' or 'female'"})
		return
	}

	if isDateValid := helpers.ValidateDateFormat(payload.BirthDate); !isDateValid {
		helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: "should be string with ISO 8601 format"})
		return
	}

	if err := helpers.ValidateURLWithDomain(payload.IdentityCardScanImg); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	//create usecase
	uu := patientusecase.New(
		patientrepository.New(u.DB),
		userrepository.New(u.DB),
	)

	// use usecase to create patient
	status, err := uu.Create(&patientusecase.ParamsCreate{
		ReqUserId:           userIdClaims,
		IdentityNumber:      payload.IdentityNumber,
		PhoneNumber:         payload.PhoneNumber,
		Name:                payload.Name,
		Gender:              strings.ToLower(payload.Gender),
		BirthDate:           payload.BirthDate,
		IdentityCardScanImg: payload.IdentityCardScanImg,
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
