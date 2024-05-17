package v1recordcontroller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dwibi/health-record/src/entities"
	"github.com/dwibi/health-record/src/helpers"
	patientrepository "github.com/dwibi/health-record/src/repository/patient"
	recordrepository "github.com/dwibi/health-record/src/repository/record"
	userrepository "github.com/dwibi/health-record/src/repository/user"
	recordusecase "github.com/dwibi/health-record/src/usecase/record"
)

func (u V1Record) FindMany(w http.ResponseWriter, r *http.Request) {
	// Access the user information from the request context
	userIdClaims, ok := r.Context().Value(helpers.UserContextKey).(int)
	if !ok {
		http.Error(w, "User information not found in context", http.StatusInternalServerError)
		return
	}

	filters := &entities.RecordSearchFilter{}
	queryParams := r.URL.Query()

	if identityStr := queryParams.Get("identityDetail.identityNumber"); identityStr != "" {
		identityNum, err := strconv.Atoi(identityStr)
		if err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{
				Message: "Invalid value for 'userId'",
			})
			return
		}
		filters.IdentityNumber = identityNum
	}

	if userIdStr := queryParams.Get("createdBy.userId"); userIdStr != "" {
		filters.UserId = userIdStr
	}

	if userNipStr := queryParams.Get("createdBy.nip"); userNipStr != "" {
		filters.UserId = userNipStr
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

	if createdAt := queryParams.Get("createdAt"); createdAt != "" {
		lowCreatedAt := strings.ToLower(createdAt)
		if lowCreatedAt == "asc" || lowCreatedAt == "desc" {
			filters.CreatedAt = lowCreatedAt
		}
	}

	fmt.Println(filters)

	// create usecase
	uu := recordusecase.New(
		recordrepository.New(u.DB),
		patientrepository.New(u.DB),
		userrepository.New(u.DB),
	)

	// use usecase to get data patient
	data, status, err := uu.FindMany(&recordusecase.ParamsFindMany{
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
