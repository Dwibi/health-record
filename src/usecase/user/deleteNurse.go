package userusecase

import (
	"errors"
	"net/http"

	"github.com/dwibi/health-record/src/helpers"
)

type ParamsDeleteUserNurse struct {
	ReqUserId int
	Id        int
}

func (i *sUserUseCase) DeleteNurse(r *ParamsUpdateUserNurseAccess) (int, error) {
	// check user role that request this call
	user, err := i.userRepository.GetUserById(r.ReqUserId)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if user == nil {
		return http.StatusUnauthorized, errors.New("unauthorize")
	}

	if isIT := helpers.IsItUser(user.NIP); !isIT {
		return http.StatusUnauthorized, errors.New("unauthorize")
	}

	// check if user exist
	user, err = i.userRepository.GetUserById(r.Id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if user == nil {
		return http.StatusNotFound, errors.New("user not found")
	}

	if isNurse := helpers.IsItNurse(user.NIP); !isNurse {
		return http.StatusNotFound, errors.New("userId isn't nurse / user nip not start with 303")
	}

	// delete nurse user into database using repository
	err = i.userRepository.Delete(r.Id)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
