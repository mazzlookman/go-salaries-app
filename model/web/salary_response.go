package web

type SalaryResponse struct {
	Id      int    `json:"id"`
	Role    string `json:"role"`
	Company string `json:"company"`
	Expr    int    `json:"expr"`
	Salary  int    `json:"salary"`
}
