package repository

import (
	"context"
	"database/sql"
	"go-salaries-app/model/domain"
)

type SalaryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, salaries domain.Salaries) domain.Salaries
	Update(ctx context.Context, tx *sql.Tx, salaries domain.Salaries) domain.Salaries
	Delete(ctx context.Context, tx *sql.Tx, salaries domain.Salaries)
	FindById(ctx context.Context, tx *sql.Tx, salaryId int) (domain.Salaries, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Salaries
}
