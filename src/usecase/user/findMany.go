package userusecase

import (
	"errors"
	"net/http"

	"github.com/dwibi/health-record/src/entities"
	"github.com/dwibi/health-record/src/helpers"
	userrepository "github.com/dwibi/health-record/src/repository/user"
)

type ParamsFindMany struct {
	QuerySearch *entities.UserSearchFilter
	ReqUserId   int
}

func (i *sUserUseCase) FindMany(r *ParamsFindMany) ([]*userrepository.ResultFindMany, int, error) {
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
	data, err := i.userRepository.FindMany(r.QuerySearch)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return data, http.StatusOK, nil
}
