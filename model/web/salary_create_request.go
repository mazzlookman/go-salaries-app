package web

type SalaryCreateRequest struct {
	Role    string
	Company string
	Expr    int
	Salary  int
}
