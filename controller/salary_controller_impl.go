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

func NewSalaryController(salaryService service.SalaryService) SalaryController {
	return &SalaryControllerImpl{SalaryService: salaryService}
}

func (controller *SalaryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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

func (controller *SalaryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	salaryUpdateRequest := web.SalaryUpdateRequest{}
	helper.ReadFromRequestBody(request, &salaryUpdateRequest)

	id := params.ByName("salaryId")
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

func (controller *SalaryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("salaryId")
	idInt, _ := strconv.Atoi(id)

	controller.SalaryService.Delete(request.Context(), idInt)

	webResp := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Salary was successfully deleted",
	}

	helper.WriteToResponseBody(writer, &webResp)
}

func (controller *SalaryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("salaryId")
	idInt, _ := strconv.Atoi(id)

	findById := controller.SalaryService.FindById(request.Context(), idInt)

	webResp := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   findById,
	}

	helper.WriteToResponseBody(writer, &webResp)
}

func (controller *SalaryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	findAll := controller.SalaryService.FindAll(request.Context())

	webResp := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   findAll,
	}

	helper.WriteToResponseBody(writer, &webResp)
}
