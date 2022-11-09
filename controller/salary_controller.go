package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type SalaryController interface {
	Create(request *http.Request, writer http.ResponseWriter, params httprouter.Params)
	Update(request *http.Request, writer http.ResponseWriter, params httprouter.Params)
	Delete(request *http.Request, writer http.ResponseWriter, params httprouter.Params)
	FindById(request *http.Request, writer http.ResponseWriter, params httprouter.Params)
	FindAll(request *http.Request, writer http.ResponseWriter, params httprouter.Params)
}
