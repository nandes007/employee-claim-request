package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CompanyController interface {
	Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Find(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}
