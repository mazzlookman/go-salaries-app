package web

type SalaryUpdateRequest struct {
	Id      int    `validate:"required"`
	Role    string `validate:"required,max=200,min=2"`
	Company string `validate:"required,max=200,min=2"`
	Expr    int    `validate:"required,max=10,min=0"`
	Salary  int    `validate:"required,max=20,min=0"`
}
