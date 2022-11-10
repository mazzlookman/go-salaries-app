package web

type SalaryCreateRequest struct {
	Role    string `validate:"required,max=200,min=2" json:"role"`
	Company string `validate:"required,max=200,min=2" json:"company"`
	Expr    int    `validate:"min=0" json:"expr"`
	Salary  int    `validate:"required" json:"salary"`
}
