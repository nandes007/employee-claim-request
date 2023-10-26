package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nandes007/employee-claim-request/helper"
	"github.com/nandes007/employee-claim-request/model/domain"
	"github.com/nandes007/employee-claim-request/model/web"
	"github.com/nandes007/employee-claim-request/model/web/auth"
	"github.com/nandes007/employee-claim-request/service"
	"net/http"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) AdminRegister(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	registerRequest := auth.RegisterRequest{}
	helper.ReadFromRequestBody(request, &registerRequest)
	apiResponse := web.ApiResponse{}

	response := controller.AuthService.Register(request.Context(), registerRequest, true)

	apiResponse.Code = 200
	apiResponse.Status = "OK"
	apiResponse.Data = response

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	registerRequest := auth.RegisterRequest{}
	helper.ReadFromRequestBody(request, &registerRequest)
	apiResponse := web.ApiResponse{}

	response := controller.AuthService.Register(request.Context(), registerRequest, false)

	apiResponse.Code = 200
	apiResponse.Status = "OK"
	apiResponse.Data = response

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *AuthControllerImpl) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	request := domain.User{}
	helper.ReadFromRequestBody(r, &request)

	token := controller.AuthService.Login(r.Context(), request)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "Success",
		Data:   token,
	}

	helper.WriteToResponseBody(w, apiResponse)
}
