package v1patientcontroller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dwibi/health-record/src/entities"
	"github.com/dwibi/health-record/src/helpers"
	patientrepository "github.com/dwibi/health-record/src/repository/patient"
	userrepository "github.com/dwibi/health-record/src/repository/user"
	patientusecase "github.com/dwibi/health-record/src/usecase/patient"
)

func (u V1Patient) FindMany(w http.ResponseWriter, r *http.Request) {
	// Access the user information from the request context
	userIdClaims, ok := r.Context().Value(helpers.UserContextKey).(int)
	if !ok {
		http.Error(w, "User information not found in context", http.StatusInternalServerError)
		return
	}

	filters := &entities.PatientSearchFilter{}
	queryParams := r.URL.Query()

	if identityStr := queryParams.Get("identityNumber"); identityStr != "" {
		identityNum, err := strconv.Atoi(identityStr)
		if err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{
				Message: "Invalid value for 'userId'",
			})
			return
		}
		filters.IdentityNumber = identityNum
	}

	if limitStr := queryParams.Get("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{
				Message: "Invalid value for 'limit'",
			})
			return
		}
		filters.Limit = limit
	}

	if offsetStr := queryParams.Get("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{
				Message: "Invalid value for 'offset'",
			})
			return
		}
		filters.Offset = offset
	}

	if name := queryParams.Get("name"); name != "" {
		filters.Name = name
	}

	if phoneStr := queryParams.Get("phoneNumber"); phoneStr != "" {
		phoneNum, err := strconv.Atoi(phoneStr)
		if err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{
				Message: "Invalid value for 'phoneNumber'",
			})
			return
		}
		filters.PhoneNumber = phoneNum
	}

	if createdAt := queryParams.Get("createdAt"); createdAt != "" {
		lowCreatedAt := strings.ToLower(createdAt)
		if lowCreatedAt == "asc" || lowCreatedAt == "desc" {
			filters.CreatedAt = lowCreatedAt
		}
	}

	fmt.Println(filters)

	// create usecase
	uu := patientusecase.New(
		patientrepository.New(u.DB),
		userrepository.New(u.DB),
	)

	// use usecase to get data patient
	data, status, err := uu.FindMany(&patientusecase.ParamsFindMany{
		ReqUserId:   userIdClaims,
		QuerySearch: filters,
	})

	if err != nil {
		helpers.WriteJSON(w, status, ErrorResponse{Message: err.Error()})
		return
	}

	helpers.WriteJSON(w, status, SuccessResponse{
		Message: "success",
		Data:    data,
	})
}
