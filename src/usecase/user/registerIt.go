package userusecase

import (
	"errors"
	"net/http"
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
	NIP         int    `json:"nip"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

func (i *sUserUseCase) RegisterIt(r *ParamsRegisterUserIt) (*ResultUser, int, error) {
	// check if nip already exists
	isNipExist, _ := i.userRepository.IsExist(r.NIP)

	if isNipExist {
		// TODO: create file for error message
		return nil, http.StatusConflict, errors.New("NIP sudah digunakan")
	}

	// Hash password
	hashedPassword, err := helpers.HashPassword(r.Password)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("failed to hash password")
	}

	// create IT user into database using repository
	data, err := i.userRepository.CreateUser(&userrepository.ParamsCreateUser{
		NIP:      r.NIP,
		Name:     r.Name,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// Generate JWT Token
	token, err := helpers.CreateUserToken(&helpers.ParamCreateUser{
		UserId:          data.ID,
		ExpiredInMinute: 400,
		SecretKey:       []byte(os.Getenv("JWT_SECRET")),
	})

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	strconv.Atoi(r.NIP)

	return &ResultUser{
		UserId:      strconv.Itoa(data.ID),
		NIP:         func() int { n, _ := strconv.Atoi(r.NIP); return n }(), // Convert r.NIP to int inline,
		Name:        r.Name,
		AccessToken: token,
	}, http.StatusCreated, nil
}
