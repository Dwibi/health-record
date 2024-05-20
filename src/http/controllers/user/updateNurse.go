package v1usercontroller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dwibi/health-record/src/helpers"
	userrepository "github.com/dwibi/health-record/src/repository/user"
	userusecase "github.com/dwibi/health-record/src/usecase/user"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type updateNurseRequest struct {
	NIP  int    `json:"nip" validate:"required,number"`
	Name string `json:"name" validate:"required,min=5,max=50"`
}

func (u V1User) UpdateNurse(w http.ResponseWriter, r *http.Request) {
	// Access the user information from the request context
	userIdClaims, ok := r.Context().Value(helpers.UserContextKey).(int)
	if !ok {
		http.Error(w, "User information not found in context", http.StatusInternalServerError)
		return
	}

	// get user nurse id
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["userId"])

	// check payload
	payload := new(updateNurseRequest)
	if err := helpers.ParseJSON(r, &payload); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	// check validation
	if err := helpers.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: fmt.Errorf("invalid payload: %v", errors).Error()})
		return
	}

	nipStr := strconv.Itoa(payload.NIP)

	// Validate NIP
	if err := helpers.ValidateNIP(nipStr); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	if isNurse := helpers.IsItNurse(nipStr); !isNurse {
		helpers.WriteJSON(w, http.StatusNotFound, ErrorResponse{Message: "nip should start 303"})
		return
	}

	// create usecase
	uu := userusecase.New(
		userrepository.New(u.DB),
	)

	// use usecase to register IT user
	status, err := uu.UpdateNurse(&userusecase.ParamsUpdateUserNurse{
		ReqUserId: userIdClaims,
		Id:        userId,
		NIP:       nipStr,
		Name:      payload.Name,
	})

	if err != nil {
		helpers.WriteJSON(w, status, ErrorResponse{Message: err.Error()})
		return
	}

	helpers.WriteJSON(w, status, SuccessResponse{
		Message: "Nurse user updated successfully",
		Data:    nil,
	})
}
