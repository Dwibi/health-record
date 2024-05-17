package patientusecase

import (
	"errors"
	"net/http"

	"github.com/dwibi/health-record/src/helpers"
	patientrepository "github.com/dwibi/health-record/src/repository/patient"
)

type ParamsCreate struct {
	ReqUserId           int
	IdentityNumber      int
	PhoneNumber         string
	Name                string
	Gender              string
	IdentityCardScanImg string
}

func (i *sPatientUseCase) Create(r *ParamsCreate) (int, error) {
	// check user role that request this call
	user, err := i.userRepository.GetUserById(r.ReqUserId)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if user == nil {
		return http.StatusUnauthorized, errors.New("unauthorize")
	}

	isIT := helpers.IsItUser(user.NIP)
	isNurse := helpers.IsItNurse(user.NIP)

	if !isIT && !isNurse {
		return http.StatusUnauthorized, errors.New("unauthorize")
	}

	// create patient
	err = i.patientRepository.Create(&patientrepository.ParamsCreate{
		IdentityNumber:      r.IdentityNumber,
		PhoneNumber:         r.PhoneNumber,
		Name:                r.Name,
		Gender:              r.Gender,
		IdentityCardScanImg: r.IdentityCardScanImg,
	})

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}
