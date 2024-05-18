package recordusecase

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dwibi/health-record/src/helpers"
	recordrepository "github.com/dwibi/health-record/src/repository/record"
)

type ParamsCreate struct {
	ReqUserId      int
	IdentityNumber int
	Symptoms       string
	Medications    string
}

func (i *sRecordUseCase) Create(r *ParamsCreate) (int, error) {
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

	// Check if the identity number already exist
	if isPatientExist, _ := i.patientRepository.IsExist(strconv.Itoa(r.IdentityNumber)); !isPatientExist {
		return http.StatusNotFound, errors.New("patient is not exists")
	}

	// create patient
	err = i.recordRepository.Create(&recordrepository.ParamsCreate{
		CreatedBy:      r.ReqUserId,
		IdentityNumber: strconv.Itoa(r.IdentityNumber),
		Symptoms:       r.Symptoms,
		Medications:    r.Medications,
	})

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}
