package controller

import (
	"github.com/julienschmidt/httprouter"
	"go-salaries-app/helper"
	"go-salaries-app/model/web"
	"go-salaries-app/service"
	"net/http"
	"strconv"
)

type SalaryControllerImpl struct {
	service.SalaryService
}

func (controller *SalaryControllerImpl) Create(request *http.Request, writer http.ResponseWriter, params httprouter.Params) {
	salaryCreateRequest := web.SalaryCreateRequest{}
	helper.ReadFromRequestBody(request, &salaryCreateRequest)

	create := controller.SalaryService.Create(request.Context(), salaryCreateRequest)

	webResp := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   create,
	}

	helper.WriteToResponseBody(writer, &webResp)

}

func (controller *SalaryControllerImpl) Update(request *http.Request, writer http.ResponseWriter, params httprouter.Params) {
	salaryUpdateRequest := web.SalaryUpdateRequest{}
	helper.ReadFromRequestBody(request, &salaryUpdateRequest)

	id := params.ByName("categoryId")
	idInt, _ := strconv.Atoi(id)

	salaryUpdateRequest.Id = idInt

	update := controller.SalaryService.Update(request.Context(), salaryUpdateRequest)

	webResp := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   update,
	}

	helper.WriteToResponseBody(writer, &webResp)
}

func (controller *SalaryControllerImpl) Delete(request *http.Request, writer http.ResponseWriter, params httprouter.Params) {
	id := params.ByName("categoryId")
	idInt, _ := strconv.Atoi(id)

	controller.SalaryService.Delete(request.Context(), idInt)

	webResp := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Salary was successfully deleted",
	}

	helper.WriteToResponseBody(writer, &webResp)
}

func (controller *SalaryControllerImpl) FindById(request *http.Request, writer http.ResponseWriter, params httprouter.Params) {
	id := params.ByName("categoryId")
	idInt, _ := strconv.Atoi(id)

	findById := controller.SalaryService.FindById(request.Context(), idInt)

	webResp := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   findById,
	}

	helper.WriteToResponseBody(writer, &webResp)
}

func (controller *SalaryControllerImpl) FindAll(request *http.Request, writer http.ResponseWriter, params httprouter.Params) {
	findAll := controller.SalaryService.FindAll(request.Context())

	webResp := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   findAll,
	}

	helper.WriteToResponseBody(writer, &webResp)
}
