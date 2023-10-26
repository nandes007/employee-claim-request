package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nandes007/employee-claim-request/helper"
	"github.com/nandes007/employee-claim-request/model/web"
	"github.com/nandes007/employee-claim-request/model/web/user"
	"github.com/nandes007/employee-claim-request/service"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{
		UserService: service,
	}
}

func (u UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	request := user.Request{}
	helper.ReadFromRequestBody(r, &request)
	token := r.Header.Get("Authorization")

	response := u.UserService.Update(r.Context(), request, token)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, apiResponse)
}
