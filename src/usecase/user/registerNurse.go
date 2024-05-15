package userusecase

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/dwibi/health-record/src/helpers"
	userrepository "github.com/dwibi/health-record/src/repository/user"
)

type ParamsRegisterUserNurse struct {
	NIP                 string
	Name                string
	IdentityCardScanImg string
}

func (i *sUserUseCase) RegisterNurse(r *ParamsRegisterUserNurse) (*ResultUser, int, error) {
	// check if nip already exists
	isNipExist, _ := i.userRepository.IsExist(r.NIP)

	if isNipExist {
		// TODO: create file for error message
		return nil, http.StatusConflict, errors.New("NIP sudah digunakan")
	}

	// create IT user into database using repository
	data, err := i.userRepository.CreateUser(&userrepository.ParamsCreateUser{
		NIP:                 r.NIP,
		Name:                r.Name,
		IdentityCardScanImg: r.IdentityCardScanImg,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// Generate JWT Token
	token, err := helpers.GenerateJwtToken(&helpers.ParamGenerateJWT{
		UserId:          data.ID,
		ExpiredInMinute: 400,
		SecretKey:       []byte(os.Getenv("JWT_SECRET")),
	})

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	strconv.Atoi(r.NIP)

	return &ResultUser{
		UserId:      strconv.FormatInt(data.ID, 10),
		NIP:         func() int { n, _ := strconv.Atoi(r.NIP); return n }(), // Convert r.NIP to int inline,
		Name:        r.Name,
		AccessToken: token,
	}, http.StatusCreated, nil
}
