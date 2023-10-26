package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ClaimRequestController interface {
	Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UpdateStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Find(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UploadFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}
