package exception

import (
	"github.com/go-playground/validator/v10"
	"github.com/nandes007/employee-claim-request/helper"
	"github.com/nandes007/employee-claim-request/model/web"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	if notFoundError(writer, request, err) {
		return
	}

	if badRequestError(writer, request, err) {
		return
	}

	if unauthorizedError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		apiResponse := web.ApiResponse{
			Code:   http.StatusUnprocessableEntity,
			Status: "Unprocessable Entity",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, apiResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		apiResponse := web.ApiResponse{
			Code:    http.StatusNotFound,
			Status:  "Not Found",
			Message: exception.Error,
		}

		helper.WriteToResponseBody(writer, apiResponse)
		return true
	} else {
		return false
	}
}

func badRequestError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(BadRequestError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		apiResponse := web.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: exception.Error,
		}

		helper.WriteToResponseBody(writer, apiResponse)
		return true
	} else {
		return false
	}
}

func unauthorizedError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(UnauthorizedError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		apiResponse := web.ApiResponse{
			Code:    http.StatusUnauthorized,
			Status:  "Unauthorized",
			Message: exception.Error,
		}

		helper.WriteToResponseBody(writer, apiResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	apiResponse := web.ApiResponse{
		Code:    http.StatusInternalServerError,
		Status:  "Internal Server Error",
		Message: "Sorry, something went wrong",
		Data:    err,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}
