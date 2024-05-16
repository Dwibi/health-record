package v1usercontroller

import (
	"net/http"
	"strconv"

	"github.com/dwibi/health-record/src/helpers"
	userrepository "github.com/dwibi/health-record/src/repository/user"
	userusecase "github.com/dwibi/health-record/src/usecase/user"
	"github.com/gorilla/mux"
)

func (u V1User) DeleteNurse(w http.ResponseWriter, r *http.Request) {
	// Access the user information from the request context
	userIdClaims, ok := r.Context().Value(helpers.UserContextKey).(int)
	if !ok {
		http.Error(w, "User information not found in context", http.StatusInternalServerError)
		return
	}

	// get user nurse id
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["userId"])

	// create usecase
	uu := userusecase.New(
		userrepository.New(u.DB),
	)

	// use usecase to register IT user
	status, err := uu.DeleteNurse(&userusecase.ParamsUpdateUserNurseAccess{
		ReqUserId: userIdClaims,
		Id:        userId,
	})

	if err != nil {
		helpers.WriteJSON(w, status, ErrorResponse{Message: err.Error()})
		return
	}

	helpers.WriteJSON(w, status, SuccessResponse{
		Message: "Update nurse access successfully",
		Data:    nil,
	})
}
