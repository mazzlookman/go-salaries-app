package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-salaries-app/helper"
	"go-salaries-app/model/domain"
)

type SalaryRepositoryImpl struct {
}

func NewSalaryRepository() SalaryRepository {
	return &SalaryRepositoryImpl{}
}

func (repository *SalaryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, salary domain.Salaries) domain.Salaries {
	sql := "insert into salaries (role, company, expr, salary) values (?,?,?,?)"
	result, err := tx.ExecContext(ctx, sql, salary.Role, salary.Company, salary.Expr, salary.Salary)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	salary.Id = int(id)
	return salary
}

func (repository *SalaryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, salary domain.Salaries) domain.Salaries {
	sql := "update salaries set role = ?, company = ?, expr = ?, salary = ? where id = ?"
	_, err := tx.ExecContext(ctx, sql, salary.Role, salary.Company, salary.Expr, salary.Salary, salary.Id)
	helper.PanicIfError(err)

	return salary
}

func (repository *SalaryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, salary domain.Salaries) {
	sql := "delete from salaries where id = ?"
	_, err := tx.ExecContext(ctx, sql, salary.Id)
	helper.PanicIfError(err)
}

func (repository *SalaryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, salaryId int) (domain.Salaries, error) {
	sql := "select * from salaries where id = ?"
	rows, err := tx.QueryContext(ctx, sql, salaryId)
	defer rows.Close()
	helper.PanicIfError(err)

	salary := domain.Salaries{}
	if rows.Next() {
		err := rows.Scan(&salary.Id, &salary.Role, &salary.Company, &salary.Expr, &salary.Salary)
		helper.PanicIfError(err)
		return salary, nil
	} else {
		return salary, errors.New("Salary is not found")
	}
}

func (repository *SalaryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Salaries {
	sql := "select * from salaries"
	rows, err := tx.QueryContext(ctx, sql)
	defer rows.Close()
	helper.PanicIfError(err)

	var salaries []domain.Salaries
	for rows.Next() {
		salary := domain.Salaries{}
		err := rows.Scan(&salary.Id, &salary.Role, &salary.Company, &salary.Expr, &salary.Salary)
		helper.PanicIfError(err)
		salaries = append(salaries, salary)
	}
	return salaries
}
