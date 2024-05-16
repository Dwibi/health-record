package userusecase

import (
	"errors"
	"net/http"

	"github.com/dwibi/health-record/src/helpers"
	userrepository "github.com/dwibi/health-record/src/repository/user"
)

type ParamsUpdateUserNurse struct {
	ReqUserId int
	Id        int
	NIP       string
	Name      string
}

func (i *sUserUseCase) UpdateNurse(r *ParamsUpdateUserNurse) (int, error) {
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

	// check if user exist and get the nip to compare
	data, err := i.userRepository.GetUserById(r.Id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if data == nil {
		return http.StatusNotFound, err

	}

	if r.NIP != data.NIP {
		isNipExist, _ := i.userRepository.IsExist(r.NIP)
		if isNipExist {
			// TODO: create file for error message
			return http.StatusConflict, errors.New("NIP sudah digunakan")
		}
	}

	// update nurse user into database using repository
	err = i.userRepository.UpdateNurse(&userrepository.ParamsUpdateNurse{
		Id:   r.Id,
		NIP:  r.NIP,
		Name: r.Name,
	})

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
