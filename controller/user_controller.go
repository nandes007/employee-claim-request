package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}
