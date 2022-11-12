package main

import (
	"github.com/go-playground/validator"
	"go-salaries-app/app"
	"go-salaries-app/controller"
	"go-salaries-app/exception"
	"go-salaries-app/helper"
	"go-salaries-app/middleware"
	"go-salaries-app/repository"
	"go-salaries-app/service"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	salaryRepository := repository.NewSalaryRepository()
	salaryService := service.NewSalaryService(salaryRepository, db, validate)
	salaryController := controller.NewSalaryController(salaryService)
	router := app.NewRouter(salaryController)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
