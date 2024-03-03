package controllers

import (
	"chi-users-project/app/utilities/httputils"
	"chi-users-project/app/services/dtos"
	"chi-users-project/app/services"
	"github.com/go-chi/chi/v5"
	"encoding/json"
	"net/http"
)


func UsersIndex(w http.ResponseWriter, r *http.Request) {
	paginationDTO := r.Context().Value("paginationDTO").(dtos.PaginationDTO)

	path := httputils.GetRequestPath(r)

	baseService := r.Context().Value("BaseService").(*services.BaseService)
	service := services.UserService{BaseService: baseService}

	result, errDTO := service.GetUsers(paginationDTO, path)
	if errDTO.Exists() {
		httputils.RenderErrorJSON(w, errDTO)
		return
	}

	httputils.RenderJSON(w, result, 200)
}


func FindUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "userId")

	baseService := r.Context().Value("BaseService").(*services.BaseService)
	service := services.UserService{BaseService: baseService}

	userOutDTO, errDTO := service.GetUser(userIdStr)
	if errDTO.Exists() {
		httputils.RenderErrorJSON(w, errDTO)
		return
	}

	httputils.RenderJSON(w, userOutDTO, 200)
}


func CreateUser(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateUserDTO
	bindErr := json.NewDecoder(r.Body).Decode(&dto)
	if bindErr != nil {
		httputils.RenderErrorJSON(w, dtos.CreateErrorDTO(bindErr, 400, false))
		return
	}

	baseService := r.Context().Value("BaseService").(*services.BaseService)
	service := services.UserService{BaseService: baseService}

	userOutDTO, errDTO := service.CreateUser(dto)

	if errDTO.Exists() {
		httputils.RenderErrorJSON(w, errDTO)
		return
	}

	httputils.RenderJSON(w, userOutDTO, 201)
}


// PATCH version of User update
// this endpoint validates the request data against the UserDTO,
// but keeps it as a map so that only the included data is updated
// (GORM only updates non-zero fields when updating with struct)
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "userId")

	var data map[string]interface{}
	bindErr := json.NewDecoder(r.Body).Decode(&data)
	if bindErr != nil {
		httputils.RenderErrorJSON(w, dtos.CreateErrorDTO(bindErr, 400, false))
		return
	}

	baseService := r.Context().Value("BaseService").(*services.BaseService)
	service := services.UserService{BaseService: baseService}

	userOutDTO, errDTO := service.UpdateUser(userIdStr, data)

	if errDTO.Exists() {
		httputils.RenderErrorJSON(w, errDTO)
		return
	}

	httputils.RenderJSON(w, userOutDTO, 200)
}


// PUT version of User update (expects all user data) (prefer above PATCH version)
func UpdateUserOG(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "userId")

	var dto dtos.UserInDTO
	bindErr := json.NewDecoder(r.Body).Decode(&dto)
	if bindErr != nil {
		httputils.RenderErrorJSON(w, dtos.CreateErrorDTO(bindErr, 400, false))
		return
	}

	baseService := r.Context().Value("BaseService").(*services.BaseService)
	service := services.UserService{BaseService: baseService}

	userOutDTO, errDTO := service.UpdateUserOG(userIdStr, dto)

	if errDTO.Exists() {
		httputils.RenderErrorJSON(w, errDTO)
		return
	}

	httputils.RenderJSON(w, userOutDTO, 200)
}


func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "userId")

	baseService := r.Context().Value("BaseService").(*services.BaseService)
	service := services.UserService{BaseService: baseService}

	errDTO := service.DeleteUser(userIdStr)

	if errDTO.Exists() {
		httputils.RenderErrorJSON(w, errDTO)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
