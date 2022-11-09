package main

import (
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"go-salaries-app/app"
	"go-salaries-app/controller"
	"go-salaries-app/repository"
	"go-salaries-app/service"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	salaryRepository := repository.NewSalaryRepository()
	salaryService := service.NewSalaryService(salaryRepository, db, validate)
	salaryController := controller.NewSalaryController(salaryService)

	router := httprouter.New()
	router.GET("/api/salaries", salaryController.FindAll)
	router.GET("/api/salaries/:salaryId", salaryController.FindById)
	router.POST("/api/salaries", salaryController.Create)
	router.PUT("/api/salaries/:salaryId", salaryController.Update)
	router.DELETE("/api/salaries/:salaryId", salaryController.Delete)
}
