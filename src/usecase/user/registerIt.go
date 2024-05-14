package userusecase

import (
	"errors"
	"os"
	"strconv"

	"github.com/dwibi/health-record/src/helpers"
	userrepository "github.com/dwibi/health-record/src/repository/user"
)

type ParamsRegisterUserIt struct {
	NIP      string
	Name     string
	Password string
}

type ResultUser struct {
	UserId      string `json:"userId"`
	NIP         string `json:"nip"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

func (i *sUserUseCase) RegisterIt(r *ParamsRegisterUserIt) (*ResultUser, error) {
	// check if nip already exists
	isNipExist, _ := i.userRepository.IsExist(r.NIP)

	if isNipExist {
		// TODO: create file for error message
		return nil, errors.New("NIP sudah digunakan")
	}

	// Hash password
	hashedPassword, err := helpers.HashPassword(r.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// create IT user into database using repository
	data, err := i.userRepository.CreateUserIt(&userrepository.ParamsCreateUserIt{
		NIP:      r.NIP,
		Name:     r.Name,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
	}

	// Generate JWT Token
	token, err := helpers.GenerateJwtToken(&helpers.ParamGenerateJWT{
		UserId:          data.ID,
		ExpiredInMinute: 400,
		SecretKey:       os.Getenv("JWT_SECRET"),
	})
	if err != nil {
		return nil, err
	}

	return &ResultUser{
		UserId:      strconv.Itoa(int(data.ID)),
		NIP:         r.NIP,
		Name:        r.Name,
		AccessToken: token,
	}, nil
}
