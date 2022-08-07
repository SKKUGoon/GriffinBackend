package rdb

type EmployeePlural struct {
	Total     int        `json:"total_length"`
	Permanent []Employee `json:"repeat_true"`
	Freelance []Employee `json:"repeat_false"`
}

type Employee struct {
	Id       int     `json:"key"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Position string  `json:"position"`
	Account  string  `json:"account"`
	Payroll  int     `json:"payroll"`
	Currency string  `json:"curr"`
	Date     int     `json:"date"`
	TimeLeft float64 `json:"secLeft"`
}

type EmployeeOnboard struct {
	MyEmployerId   string
	MyEmployeeId   int
	EmploymentType string
}

type Payment struct {
	Name     string `json:"name"`
	Payroll  int    `json:"payroll"`
	Currency string `json:"curr"`
	Time     string `json:"time"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
