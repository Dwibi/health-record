package userusecase

import (
	userrepository "github.com/dwibi/health-record/src/repository/user"
)

type sUserUseCase struct {
	userRepository userrepository.IUserRepository
}

type IUserUseCase interface {
	RegisterIt(*ParamsRegisterUserIt) (*ResultUser, int, error)
	LoginIt(*ParamsLoginUserIt) (*ResultUser, int, error)
	LoginNurse(r *ParamsLoginUserIt) (*ResultUser, int, error)
	RegisterNurse(r *ParamsRegisterUserNurse) (*ResultUser, int, error)
	UpdateNurse(r *ParamsUpdateUserNurse) (int, error)
	UpdateNurseAccess(r *ParamsUpdateUserNurseAccess) (int, error)
	DeleteNurse(r *ParamsUpdateUserNurseAccess) (int, error)
	FindMany(r *ParamsFindMany) ([]*userrepository.ResultFindMany, int, error)
}

func New(userRepository userrepository.IUserRepository) IUserUseCase {
	return &sUserUseCase{
		userRepository: userRepository,
	}
}
