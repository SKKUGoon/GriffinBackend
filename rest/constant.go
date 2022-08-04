package rest

const (
	// function get employer query
	EMPLOYER_ID     = "employerId"
	EMPLOYMENT_TYPE = "employType"
	EMPLOYEE_ID     = "id"
	// function new employer query
	EMPLOYEE_NAME     = "name"
	EMPLOYEE_EMAIL    = "email"
	EMPLOYEE_POSITION = "position"
	EMPLOYEE_ACCOUNT  = "account"
	EMPLOYEE_PAYROLL  = "payroll"
	EMPLOYEE_PAYDAY   = "date"
	EMPLOYEE_CURRENCY = "curr"
	// function new employer query index
	EMPLOYEE_NAME_INDEX     = 0
	EMPLOYEE_EMAIL_INDEX    = 1
	EMPLOYEE_POSITION_INDEX = 2
	EMPLOYEE_ACCOUNT_INDEX  = 3
	EMPLOYEE_PAYROLL_INDEX  = 4
	EMPLOYEE_PAYDAY_INDEX   = 5
	EMPLOYEE_CURRENCY_INDEX = 6
	// database access keys
	PERMANENT_EMPLOYER_KEY  = "employee_perma:%v"
	PERMANENT_EMPLOYER_PATH = "$"
	FREELANCE_EMPLOYER_KEY  = "employee_free:%v"
	FREELANCE_EMPLOYER_PATH = "$"
	// database message success & error
	DATABASE_APPEND_SUCCESS = "database new data added"
	DATABASE_APPEND_FAIL    = "fail to append new data"
	DATABASE_CREATE_SUCCESS = "database new key-value pair added"
	DATABASE_CREATE_FAIL    = "fail to create new key-value pair"
	DATABASE_DELETE_SUCCESS = "database requested data deleted"
	DATABASE_DELETE_FAIL    = "fail to delete requested data"
	// request messages
	REQUEST_WRONG_TYPE    = "wrong type parameter"
	REQUEST_MISSING_PARAM = "missing parameter"
)

var EmployTypeMap = map[string]string{
	"perma_key":  PERMANENT_EMPLOYER_KEY,
	"perma_path": PERMANENT_EMPLOYER_PATH,
	"free_key":   FREELANCE_EMPLOYER_KEY,
	"free_path":  FREELANCE_EMPLOYER_PATH,
}
