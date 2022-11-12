package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator"
	"go-salaries-app/app"
	"go-salaries-app/controller"
	"go-salaries-app/helper"
	"go-salaries-app/middleware"
	"go-salaries-app/model/domain"
	"go-salaries-app/repository"
	"go-salaries-app/service"
	"gopkg.in/go-playground/assert.v1"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TruncateSalary(db *sql.DB) {
	db.Exec("truncate table salaries")
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	salaryRepository := repository.NewSalaryRepository()
	salaryService := service.NewSalaryService(salaryRepository, db, validate)
	salaryController := controller.NewSalaryController(salaryService)
	router := app.NewRouter(salaryController)

	return middleware.NewAuthMiddleware(router)
}

func TestCreateSalarySuccess(t *testing.T) {
	db := app.NewDBTest()
	TruncateSalary(db)

	router := setupRouter(db)

	payload := strings.NewReader(`{"role":"UI/UX","company":"Bukalapak","expr":2,"salary":10000000}`)

	request := httptest.NewRequest("POST", "http://localhost:8080/api/salaries", payload)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-api-key", "rahasia")

	writer := httptest.NewRecorder()

	router.ServeHTTP(writer, request)

	result := writer.Result()

	bytes, _ := io.ReadAll(result.Body)
	var salary map[string]interface{}
	json.Unmarshal(bytes, &salary)

	assert.Equal(t, 200, int(salary["code"].(float64)))
	assert.Equal(t, "OK", salary["status"])
	assert.Equal(t, "UI/UX", salary["data"].(map[string]interface{})["role"])
}

func TestCreateSalaryFailed(t *testing.T) {
	db := app.NewDBTest()
	TruncateSalary(db)

	router := setupRouter(db)

	payload := strings.NewReader(`{"role":"","company":"","expr":null,"salary":null}`)

	request := httptest.NewRequest("POST", "http://localhost:8080/api/salaries", payload)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-api-key", "rahasia")

	writer := httptest.NewRecorder()

	router.ServeHTTP(writer, request)

	result := writer.Result()

	bytes, _ := io.ReadAll(result.Body)
	var salary map[string]interface{}
	json.Unmarshal(bytes, &salary)

	assert.Equal(t, 400, int(salary["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", salary["status"])
}

func TestUpdateSalarySuccess(t *testing.T) {
	db := app.NewDBTest()
	TruncateSalary(db)
	router := setupRouter(db)

	//create new salary
	salaryRepository := repository.NewSalaryRepository()
	tx, err := db.Begin()
	save := salaryRepository.Save(context.Background(), tx, domain.Salaries{
		Role:    "Technical Architect",
		Company: "Blibli",
		Expr:    10,
		Salary:  50000000,
	})
	helper.PanicIfError(err)
	tx.Commit()

	//update salary that was created above
	payload := strings.NewReader(`{"role":"CTO","company":"Bukalapak","expr":15,"salary":100000000}`)
	request := httptest.NewRequest("PUT", "http://localhost:8080/api/salaries/"+strconv.Itoa(save.Id), payload)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-api-key", "rahasia")
	writer := httptest.NewRecorder()

	router.ServeHTTP(writer, request)

	response := writer.Result()
	bytes, _ := io.ReadAll(response.Body)

	var salary map[string]interface{}
	json.Unmarshal(bytes, &salary)

	assert.Equal(t, 200, int(salary["code"].(float64)))
	assert.Equal(t, "OK", salary["status"])
	assert.Equal(t, save.Id, int(salary["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "CTO", salary["data"].(map[string]interface{})["role"])
}

func TestUpdateSalaryFailed(t *testing.T) {
	db := app.NewDBTest()
	TruncateSalary(db)
	router := setupRouter(db)

	//create new salary
	salaryRepository := repository.NewSalaryRepository()
	tx, err := db.Begin()
	save := salaryRepository.Save(context.Background(), tx, domain.Salaries{
		Role:    "Technical Architect",
		Company: "Blibli",
		Expr:    10,
		Salary:  50000000,
	})
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//update salary that was created above
	payload := strings.NewReader(`{"role":"","company":"","expr":null,"salary":null}`)
	request := httptest.NewRequest("PUT", "http://localhost:8080/api/salaries/"+strconv.Itoa(save.Id), payload)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-api-key", "rahasia")
	writer := httptest.NewRecorder()

	router.ServeHTTP(writer, request)

	response := writer.Result()
	bytes, _ := io.ReadAll(response.Body)

	var salary map[string]interface{}
	json.Unmarshal(bytes, &salary)

	assert.Equal(t, 400, int(salary["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", salary["status"])
	//assert.Equal(t, save.Id, int(salary["data"].(map[string]interface{})["id"].(float64)))
	//assert.Equal(t, "CTO", salary["data"].(map[string]interface{})["role"])
}

func TestGetSalarySuccess(t *testing.T) {

}

func TestGetSalaryFailed(t *testing.T) {

}

func TestDeleteSalarySuccess(t *testing.T) {

}

func TestDeleteSalaryFailed(t *testing.T) {

}

func TestListSalariesSuccess(t *testing.T) {

}

func TestUnauthorized(t *testing.T) {

}