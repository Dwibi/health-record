package patientusecase

import (
	patientrepository "github.com/dwibi/health-record/src/repository/patient"
	userrepository "github.com/dwibi/health-record/src/repository/user"
)

type sPatientUseCase struct {
	patientRepository patientrepository.IPatientRepository
	userRepository    userrepository.IUserRepository
}

type IPatientUseCase interface {
	// RegisterIt(*ParamsRegisterUserIt) (*ResultUser, int, error)
	Create(*ParamsCreate) (int, error)
	FindMany(r *ParamsFindMany) ([]*patientrepository.ResultFindMany, int, error)
	// FindMany(r *ParamsFindMany) ([]*userrepository.ResultFindMany, int, error)
}

func New(patientRepository patientrepository.IPatientRepository, userRepository userrepository.IUserRepository) IPatientUseCase {
	return &sPatientUseCase{
		patientRepository: patientRepository,
		userRepository:    userRepository,
	}
}
