package uploadusecase

import (
	userrepository "github.com/dwibi/health-record/src/repository/user"
)

type sUploadUseCase struct {
	userRepository userrepository.IUserRepository
}

type IUploadUseCase interface {
	ValidateRole(int) (int, error)
}

func New(userRepository userrepository.IUserRepository) IUploadUseCase {
	return &sUploadUseCase{
		userRepository: userRepository,
	}
}
