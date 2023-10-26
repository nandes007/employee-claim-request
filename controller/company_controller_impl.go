package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/nandes007/employee-claim-request/exception"
	"github.com/nandes007/employee-claim-request/helper"
	"github.com/nandes007/employee-claim-request/model/web"
	"github.com/nandes007/employee-claim-request/model/web/company"
	"github.com/nandes007/employee-claim-request/service"
)

type CompanyControllerImpl struct {
	CompanyService service.CompanyService
}

func NewCompanyController(service service.CompanyService) CompanyController {
	return &CompanyControllerImpl{
		CompanyService: service,
	}
}

func (c CompanyControllerImpl) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	request := company.Request{}
	helper.ReadFromRequestBody(r, &request)

	response := c.CompanyService.Create(r.Context(), request)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, apiResponse)
}

func (c CompanyControllerImpl) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	request := company.Request{}
	helper.ReadFromRequestBody(r, &request)

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		err = errors.New("sorry, claim company has no longer exist")
		panic(exception.NewNotFoundError(err.Error()))
	}

	response := c.CompanyService.Update(r.Context(), request, id)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, apiResponse)
}

func (c CompanyControllerImpl) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		err = errors.New("sorry, claim company has no longer exist")
		panic(exception.NewNotFoundError(err.Error()))
	}

	c.CompanyService.Delete(r.Context(), id)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, apiResponse)
}

func (c CompanyControllerImpl) Find(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		err = errors.New("sorry, claim company has no longer exist")
		panic(exception.NewNotFoundError(err.Error()))
	}

	response := c.CompanyService.Find(r.Context(), id)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, apiResponse)
}

func (c CompanyControllerImpl) GetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	responses := c.CompanyService.GetAll(r.Context())
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	}

	helper.WriteToResponseBody(w, apiResponse)
}
