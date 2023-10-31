package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/nandes007/employee-claim-request/exception"
	"github.com/nandes007/employee-claim-request/helper"
	"github.com/nandes007/employee-claim-request/model/web"
	"github.com/nandes007/employee-claim-request/model/web/claim_request"
	"github.com/nandes007/employee-claim-request/service"
)

type ClaimRequestControllerImpl struct {
	ClaimRequestService service.ClaimRequestService
}

func NewClaimRequestController(service service.ClaimRequestService) ClaimRequestController {
	return &ClaimRequestControllerImpl{
		ClaimRequestService: service,
	}
}

func (c ClaimRequestControllerImpl) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	request := claim_request.Request{}
	helper.ReadFromRequestBody(r, &request)
	token := r.Header.Get("Authorization")

	response := c.ClaimRequestService.Create(r.Context(), request, token)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, apiResponse)
}

func (c ClaimRequestControllerImpl) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	request := claim_request.Request{}
	helper.ReadFromRequestBody(r, &request)
	token := r.Header.Get("Authorization")

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		err = errors.New("sorry, claim request has no longer exist")
		panic(exception.NewNotFoundError(err.Error()))
	}

	response := c.ClaimRequestService.Update(r.Context(), request, id, token)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, apiResponse)
}

func (c ClaimRequestControllerImpl) UpdateStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	request := claim_request.ApproveRequest{}
	helper.ReadFromRequestBody(r, &request)
	token := r.Header.Get("Authorization")

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		err = errors.New("sorry, claim request has no longer exist")
		panic(exception.NewNotFoundError(err.Error()))
	}

	response := c.ClaimRequestService.UpdateStatus(r.Context(), request, id, token)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, apiResponse)
}

func (c ClaimRequestControllerImpl) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		err = errors.New("sorry, claim request has no longer exist")
		panic(exception.NewNotFoundError(err.Error()))
	}

	token := r.Header.Get("Authorization")
	c.ClaimRequestService.Delete(r.Context(), id, token)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, apiResponse)
}

func (c ClaimRequestControllerImpl) Find(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		err = errors.New("sorry, claim request has no longer exist")
		panic(exception.NewNotFoundError(err.Error()))
	}

	token := r.Header.Get("Authorization")
	response := c.ClaimRequestService.Find(r.Context(), id, token)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, apiResponse)
}

func (c ClaimRequestControllerImpl) GetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token := r.Header.Get("Authorization")
	responses := c.ClaimRequestService.GetAll(r.Context(), token)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	}

	helper.WriteToResponseBody(w, apiResponse)
}
