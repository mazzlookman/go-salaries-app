package web

type SalaryUpdateRequest struct {
	Id      int
	Role    string
	Company string
	Expr    int
	Salary  int
}
