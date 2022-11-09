package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"go-salaries-app/helper"
	"go-salaries-app/model/domain"
	"go-salaries-app/model/web"
	"go-salaries-app/repository"
)

type SalaryServiceImpl struct {
	repository.SalaryRepository
	*sql.DB
	*validator.Validate
}

func NewSalaryService(salaryRepository repository.SalaryRepository, DB *sql.DB, validate *validator.Validate) SalaryService {
	return &SalaryServiceImpl{
		SalaryRepository: salaryRepository,
		DB:               DB,
		Validate:         validate}
}

func (service *SalaryServiceImpl) Create(ctx context.Context, request web.SalaryCreateRequest) web.SalaryResponse {
	errValidate := service.Validate.Struct(request)
	helper.PanicIfError(errValidate)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	salary := domain.Salaries{
		Role:    request.Role,
		Company: request.Company,
		Expr:    request.Expr,
		Salary:  request.Salary,
	}

	saveResponse := service.SalaryRepository.Save(ctx, tx, salary)

	return helper.ToSalaryResponse(saveResponse)
}

func (service *SalaryServiceImpl) Update(ctx context.Context, request web.SalaryUpdateRequest) web.SalaryResponse {
	errValidate := service.Validate.Struct(request)
	helper.PanicIfError(errValidate)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	findByIdResponse, err := service.SalaryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	findByIdResponse.Role = request.Role
	findByIdResponse.Company = request.Company
	findByIdResponse.Expr = request.Expr
	findByIdResponse.Salary = request.Salary

	updateResponse := service.SalaryRepository.Update(ctx, tx, findByIdResponse)

	return helper.ToSalaryResponse(updateResponse)
}

func (service *SalaryServiceImpl) Delete(ctx context.Context, salaryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	findByIdResponse, err := service.SalaryRepository.FindById(ctx, tx, salaryId)
	helper.PanicIfError(err)

	service.SalaryRepository.Delete(ctx, tx, findByIdResponse)
}

func (service *SalaryServiceImpl) FindById(ctx context.Context, salaryId int) web.SalaryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	findByIdResponse, err := service.SalaryRepository.FindById(ctx, tx, salaryId)
	helper.PanicIfError(err)

	return helper.ToSalaryResponse(findByIdResponse)
}

func (service *SalaryServiceImpl) FindAll(ctx context.Context) []web.SalaryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	findAllResponse := service.SalaryRepository.FindAll(ctx, tx)

	return helper.ToSalariesResponse(findAllResponse)
}
