package recordusecase

import (
	"errors"
	"net/http"

	"github.com/dwibi/health-record/src/entities"
	"github.com/dwibi/health-record/src/helpers"
	recordrepository "github.com/dwibi/health-record/src/repository/record"
)

type ParamsFindMany struct {
	QuerySearch *entities.RecordSearchFilter
	ReqUserId   int
}

func (i *sRecordUseCase) FindMany(r *ParamsFindMany) ([]*recordrepository.ResultFindMany, int, error) {
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
	data, err := i.recordRepository.FindMany(r.QuerySearch)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return data, http.StatusOK, nil
}
