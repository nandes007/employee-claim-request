package controller

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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

func (c ClaimRequestControllerImpl) UploadFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse the multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("ERROR1")
		return
	}

	// Retrieve the file from the request
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("ERROR2")
		return
	}
	defer file.Close()

	// Specify the directory where you want to save the uploaded files
	uploadDir := "uploads/"

	// Create the directory if it doesn't exist
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.MkdirAll(uploadDir, os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Create a new file on the server and copy the uploaded file to it
	newFile, err := os.Create("uploads/" + handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("ERROR3")
		return
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("ERROR4")
		return
	}

	// Get the server address from the request
	serverAddr := fmt.Sprintf("%s://%s", getProto(r), r.Host)

	// Construct the URL for the uploaded file
	fileURL := fmt.Sprintf("%s/uploads/%s", serverAddr, handler.Filename)

	// Return the URL as a response
	fmt.Fprintf(w, "File uploaded successfully. You can access it at: <a href='%s'>%s</a>", fileURL, fileURL)
}

func getProto(r *http.Request) string {
	if r.TLS != nil {
		return "https"
	}
	return "http"
}
