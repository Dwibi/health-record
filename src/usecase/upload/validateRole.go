package uploadusecase

import (
	"errors"
	"net/http"

	"github.com/dwibi/health-record/src/helpers"
)

type ParamsCreate struct {
	ReqUserId           int
	IdentityNumber      int
	PhoneNumber         string
	Name                string
	Gender              string
	BirthDate           string
	IdentityCardScanImg string
}

func (i *sUploadUseCase) ValidateRole(reqUserId int) (int, error) {
	// check user role that request this call
	user, err := i.userRepository.GetUserById(reqUserId)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if user == nil {
		return http.StatusUnauthorized, errors.New("unauthorize")
	}

	isIT := helpers.IsItUser(user.NIP)
	isNurse := helpers.IsItNurse(user.NIP)

	if !isIT && !isNurse {
		return http.StatusUnauthorized, errors.New("unauthorize")
	}

	return http.StatusOK, nil
}
