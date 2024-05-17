package patientusecase

import (
	"errors"
	"net/http"

	"github.com/dwibi/health-record/src/entities"
	"github.com/dwibi/health-record/src/helpers"
	patientrepository "github.com/dwibi/health-record/src/repository/patient"
)

type ParamsFindMany struct {
	QuerySearch *entities.PatientSearchFilter
	ReqUserId   int
}

func (i *sPatientUseCase) FindMany(r *ParamsFindMany) ([]*patientrepository.ResultFindMany, int, error) {
	// check user role that request this call
	user, err := i.userRepository.GetUserById(r.ReqUserId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if user == nil {
		return nil, http.StatusUnauthorized, errors.New("unauthorize")
	}

	if isIT := helpers.IsItUser(user.NIP); !isIT {
		return nil, http.StatusUnauthorized, errors.New("unauthorize")
	}

	// Get data from repository
	data, err := i.patientRepository.FindMany(r.QuerySearch)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return data, http.StatusOK, nil
}
