package test

import (
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator"
	"go-salaries-app/app"
	"go-salaries-app/controller"
	"go-salaries-app/middleware"
	"go-salaries-app/repository"
	"go-salaries-app/service"
	"gopkg.in/go-playground/assert.v1"
	"io"
	"net/http"
	"net/http/httptest"
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

	requestBody := strings.NewReader(`{"role":"UI/UX","company":"Bukalapak","expr":2,"salary":10000000}`)

	request := httptest.NewRequest("POST", "http://localhost:8080/api/salaries", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-api-key", "rahasia")

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	result := response.Result()
	assert.Equal(t, 200, result.StatusCode)

	bytes, _ := io.ReadAll(result.Body)
	var salary map[string]interface{}
	json.Unmarshal(bytes, &salary)

	assert.Equal(t, 200, int(salary["code"].(float64)))
	assert.Equal(t, "OK", salary["status"])
	assert.Equal(t, "UI/UX", salary["data"].(map[string]interface{})["role"])

}

func TestCreateSalaryFailed(t *testing.T) {

}

func TestUpdateSalarySuccess(t *testing.T) {

}

func TestUpdateSalaryFailed(t *testing.T) {

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
