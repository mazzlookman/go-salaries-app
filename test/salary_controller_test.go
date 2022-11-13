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
	db := app.NewDBTest()
	TruncateSalary(db)
	router := setupRouter(db)

	tx, err := db.Begin()
	helper.PanicIfError(err)

	salaryRepository := repository.NewSalaryRepository()
	save := salaryRepository.Save(context.Background(), tx, domain.Salaries{
		Role:    "Software Engineer",
		Company: "LinkAja",
		Expr:    10,
		Salary:  50000000,
	})

	helper.CommitOrRollback(tx)

	request := httptest.NewRequest("GET", "http://localhost:8080/api/salaries/"+strconv.Itoa(save.Id), nil)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("x-api-key", "rahasia")

	writer := httptest.NewRecorder()

	router.ServeHTTP(writer, request)

	response := writer.Result()
	bytes, _ := io.ReadAll(response.Body)

	var salary map[string]interface{}
	json.Unmarshal(bytes, &salary)

	assert.Equal(t, 200, int(salary["code"].(float64)))
	assert.Equal(t, "OK", salary["status"])
	assert.Equal(t, save.Role, salary["data"].(map[string]interface{})["role"])
	assert.Equal(t, save.Id, int(salary["data"].(map[string]interface{})["id"].(float64)))
}

func TestGetSalaryFailed(t *testing.T) {
	db := app.NewDBTest()
	TruncateSalary(db)
	router := setupRouter(db)

	tx, err := db.Begin()
	helper.PanicIfError(err)

	salaryRepository := repository.NewSalaryRepository()
	salaryRepository.Save(context.Background(), tx, domain.Salaries{
		Role:    "Software Engineer",
		Company: "LinkAja",
		Expr:    10,
		Salary:  50000000,
	})

	helper.CommitOrRollback(tx)

	request := httptest.NewRequest("GET", "http://localhost:8080/api/salaries/404", nil)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("x-api-key", "rahasia")

	writer := httptest.NewRecorder()

	router.ServeHTTP(writer, request)

	response := writer.Result()
	bytes, _ := io.ReadAll(response.Body)

	var salary map[string]interface{}
	json.Unmarshal(bytes, &salary)

	assert.Equal(t, 404, int(salary["code"].(float64)))
	assert.Equal(t, "NOT FOUND", salary["status"])
}

func TestDeleteSalarySuccess(t *testing.T) {
	db := app.NewDBTest()
	TruncateSalary(db)
	router := setupRouter(db)

	tx, err := db.Begin()
	helper.PanicIfError(err)

	salaryRepository := repository.NewSalaryRepository()
	save := salaryRepository.Save(context.Background(), tx, domain.Salaries{
		Role:    "Software Engineer",
		Company: "LinkAja",
		Expr:    10,
		Salary:  50000000,
	})

	helper.CommitOrRollback(tx)

	request := httptest.NewRequest("DELETE", "http://localhost:8080/api/salaries/"+strconv.Itoa(save.Id), nil)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("x-api-key", "rahasia")

	writer := httptest.NewRecorder()

	router.ServeHTTP(writer, request)

	response := writer.Result()
	bytes, _ := io.ReadAll(response.Body)

	var salary map[string]interface{}
	json.Unmarshal(bytes, &salary)

	assert.Equal(t, 200, int(salary["code"].(float64)))
	assert.Equal(t, "OK", salary["status"])
	assert.Equal(t, "Salary was successfully deleted", salary["data"])
}

func TestDeleteSalaryFailed(t *testing.T) {
	db := app.NewDBTest()
	TruncateSalary(db)
	router := setupRouter(db)

	tx, err := db.Begin()
	helper.PanicIfError(err)

	salaryRepository := repository.NewSalaryRepository()
	salaryRepository.Save(context.Background(), tx, domain.Salaries{
		Role:    "Software Engineer",
		Company: "LinkAja",
		Expr:    10,
		Salary:  50000000,
	})

	helper.CommitOrRollback(tx)

	request := httptest.NewRequest("DELETE", "http://localhost:8080/api/salaries/404", nil)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("x-api-key", "rahasia")

	writer := httptest.NewRecorder()

	router.ServeHTTP(writer, request)

	response := writer.Result()
	bytes, _ := io.ReadAll(response.Body)

	var salary map[string]interface{}
	json.Unmarshal(bytes, &salary)

	assert.Equal(t, 404, int(salary["code"].(float64)))
	assert.Equal(t, "NOT FOUND", salary["status"])
}

func TestListSalariesSuccess(t *testing.T) {
	db := app.NewDBTest()
	TruncateSalary(db)
	router := setupRouter(db)

	tx, err := db.Begin()
	helper.PanicIfError(err)

	salaryRepository := repository.NewSalaryRepository()
	save1 := salaryRepository.Save(context.Background(), tx, domain.Salaries{
		Role:    "Software Engineer",
		Company: "LinkAja",
		Expr:    10,
		Salary:  50000000,
	})

	save2 := salaryRepository.Save(context.Background(), tx, domain.Salaries{
		Role:    "Data Analyst",
		Company: "LinkAja",
		Expr:    5,
		Salary:  25000000,
	})

	helper.CommitOrRollback(tx)

	request := httptest.NewRequest("GET", "http://localhost:8080/api/salaries", nil)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("x-api-key", "rahasia")

	writer := httptest.NewRecorder()

	router.ServeHTTP(writer, request)

	response := writer.Result()
	bytes, _ := io.ReadAll(response.Body)

	var salaries map[string]interface{}
	json.Unmarshal(bytes, &salaries)

	var getAll = salaries["data"].([]interface{})
	salary1 := getAll[0].(map[string]interface{})
	salary2 := getAll[1].(map[string]interface{})

	assert.Equal(t, 200, int(salaries["code"].(float64)))
	assert.Equal(t, "OK", salaries["status"])
	assert.Equal(t, save1.Role, salary1["role"])
	assert.Equal(t, save2.Role, salary2["role"])

}

func TestUnauthorized(t *testing.T) {

}
