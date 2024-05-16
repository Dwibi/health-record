package userusecase

import (
	"errors"
	"log"
	"net/http"

	"github.com/dwibi/health-record/src/helpers"
	userrepository "github.com/dwibi/health-record/src/repository/user"
)

type ParamsUpdateUserNurseAccess struct {
	ReqUserId int
	Id        int
	Password  string
}

func (i *sUserUseCase) UpdateNurseAccess(r *ParamsUpdateUserNurseAccess) (int, error) {
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
		log.Println("masuk ke error 404")
		return http.StatusNotFound, errors.New("user not found")
	}

	if isNurse := helpers.IsItNurse(user.NIP); !isNurse {
		return http.StatusBadRequest, errors.New("can only give access to nurse")
	}

	// Hash password
	hashedPassword, err := helpers.HashPassword(r.Password)
	if err != nil {
		return http.StatusInternalServerError, errors.New("failed to hash password")
	}

	// update nurse user into database using repository
	err = i.userRepository.UpdatePassword(&userrepository.ParamsUpdatePassword{
		Id:       r.Id,
		Password: hashedPassword,
	})

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
