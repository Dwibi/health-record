package v1usercontroller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dwibi/health-record/src/entities"
	"github.com/dwibi/health-record/src/helpers"
	userrepository "github.com/dwibi/health-record/src/repository/user"
	userusecase "github.com/dwibi/health-record/src/usecase/user"
)

func (u V1User) FindMany(w http.ResponseWriter, r *http.Request) {
	// Access the user information from the request context
	userIdClaims, ok := r.Context().Value(helpers.UserContextKey).(int)
	if !ok {
		http.Error(w, "User information not found in context", http.StatusInternalServerError)
		return
	}

	filters := &entities.UserSearchFilter{}
	queryParams := r.URL.Query()

	if userIdStr := queryParams.Get("userId"); userIdStr != "" {
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{
				Message: "Invalid value for 'userId'",
			})
			return
		}
		filters.UserId = userId
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

	if nipStr := queryParams.Get("nip"); nipStr != "" {
		nip, err := strconv.Atoi(nipStr)
		if err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, ErrorResponse{
				Message: "Invalid value for 'nip'",
			})
			return
		}
		filters.NIP = nip
	}

	if role := queryParams.Get("role"); role != "" {
		filters.Role = role
	}

	if createdAt := queryParams.Get("createdAt"); createdAt != "" {
		lowCreatedAt := strings.ToLower(createdAt)
		if lowCreatedAt == "asc" || lowCreatedAt == "desc" {
			filters.CreatedAt = lowCreatedAt
		}
	}

	// fmt.Println("||||||||||||")
	// fmt.Println(filters)

	// create usecase
	uu := userusecase.New(
		userrepository.New(u.DB),
	)

	// use usecase to register IT user
	data, status, err := uu.FindMany(&userusecase.ParamsFindMany{
		QuerySearch: filters,
		ReqUserId:   userIdClaims,
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
