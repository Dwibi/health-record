package recordusecase

import (
	patientrepository "github.com/dwibi/health-record/src/repository/patient"
	recordrepository "github.com/dwibi/health-record/src/repository/record"
	userrepository "github.com/dwibi/health-record/src/repository/user"
)

type sRecordUseCase struct {
	recordRepository  recordrepository.IRecordRepository
	patientRepository patientrepository.IPatientRepository
	userRepository    userrepository.IUserRepository
}

type IRecordUseCase interface {
	// RegisterIt(*ParamsRegisterUserIt) (*ResultUser, int, error)
	Create(*ParamsCreate) (int, error)
	FindMany(r *ParamsFindMany) ([]*recordrepository.ResultFindMany, int, error)
	// FindMany(r *ParamsFindMany) ([]*patientrepository.ResultFindMany, int, error)
	// FindMany(r *ParamsFindMany) ([]*userrepository.ResultFindMany, int, error)
}

func New(recordRepository recordrepository.IRecordRepository, patientRepository patientrepository.IPatientRepository, userRepository userrepository.IUserRepository) IRecordUseCase {
	return &sRecordUseCase{
		recordRepository:  recordRepository,
		patientRepository: patientRepository,
		userRepository:    userRepository,
	}
}
