package service

import (
	"context"
	"go-salaries-app/model/web"
)

type SalaryService interface {
	Create(ctx context.Context, request web.SalaryCreateRequest) web.SalaryResponse
	Update(ctx context.Context, request web.SalaryUpdateRequest) web.SalaryResponse
	Delete(ctx context.Context, salaryId int)
	FindById(ctx context.Context, salaryId int) web.SalaryResponse
	FindAll(ctx context.Context) []web.SalaryResponse
}
