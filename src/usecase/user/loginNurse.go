package userusecase

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/dwibi/health-record/src/helpers"
	userrepository "github.com/dwibi/health-record/src/repository/user"
)

func (i *sUserUseCase) LoginNurse(r *ParamsLoginUserIt) (*ResultUser, int, error) {
	// Find IT user into database using repository
	data, err := i.userRepository.FindOneUser(&userrepository.ParamsFindOneUser{
		NIP: r.NIP,
	})

	// log.Println(data)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if data == nil {
		return nil, http.StatusNotFound, errors.New("user not found")
	}

	if data.Password == "" {
		return nil, http.StatusBadRequest, errors.New("user is not having access")
	}

	// Compare requested password with hashed password
	if isPasswordSame := helpers.ComparePassword(r.Password, data.Password); !isPasswordSame {
		return nil, http.StatusBadRequest, errors.New("password is wrong")
	}

	nipInt, _ := strconv.Atoi(r.NIP)

	// Generate JWT Token
	token, err := helpers.CreateUserToken(&helpers.ParamCreateUser{
		UserId:          data.ID,
		ExpiredInMinute: 400,
		SecretKey:       []byte(os.Getenv("JWT_SECRET")),
	})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &ResultUser{
		UserId:      strconv.Itoa(data.ID),
		NIP:         nipInt,
		Name:        data.Name,
		AccessToken: token,
	}, http.StatusOK, nil
}
