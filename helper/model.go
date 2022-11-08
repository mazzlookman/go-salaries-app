package helper

import (
	"go-salaries-app/model/domain"
	"go-salaries-app/model/web"
)

func ToSalaryResponse(salaries domain.Salaries) web.SalaryResponse {
	salary := web.SalaryResponse{
		Id:      salaries.Id,
		Role:    salaries.Role,
		Company: salaries.Company,
		Expr:    salaries.Expr,
		Salary:  salaries.Salary,
	}
	return salary
}

func ToSalariesResponse(salaries []domain.Salaries) []web.SalaryResponse {
	var salariesResp []web.SalaryResponse
	for _, salary := range salaries {
		salaryResp := web.SalaryResponse{}
		salaryResp.Id = salary.Id
		salaryResp.Role = salary.Role
		salaryResp.Company = salary.Company
		salaryResp.Expr = salary.Expr
		salaryResp.Salary = salary.Salary

		salariesResp = append(salariesResp, salaryResp)
	}
	return salariesResp
}
