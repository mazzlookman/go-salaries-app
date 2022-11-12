package app

import (
	"github.com/julienschmidt/httprouter"
	"go-salaries-app/controller"
)

func NewRouter(salaryController controller.SalaryController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/salaries", salaryController.FindAll)
	router.GET("/api/salaries/:salaryId", salaryController.FindById)
	router.POST("/api/salaries", salaryController.Create)
	router.PUT("/api/salaries/:salaryId", salaryController.Update)
	router.DELETE("/api/salaries/:salaryId", salaryController.Delete)

	return router
}
