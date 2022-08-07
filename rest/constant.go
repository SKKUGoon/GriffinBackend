package rest

const (
	// access control
	ALLOW_ORIGIN            = "Access-Control-Allow-Origin"
	ALLOW_ORIGIN_VALUE      = "*"
	ALLOW_CREDENTIALS       = "Access-Control-Allow-Credentials"
	ALLOW_CREDENTIALS_VALUE = "true"
	// function get employer query
	EMPLOYER_ID      = "employerId"
	EMPLOYER_PW      = "employerPw"
	EMPLOYMENT_TYPE  = "employType"
	EMPLOYEE_ID      = "key"
	EMPLOYEE_PARTIAL = "isPartial"
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
	HISTORICAL_PAYMENT_KEY  = "hist_payroll:%v"
	HISTORICAL_PAYMENT_PATH = "$"
	LOGIN_KEY               = "login:0"
	LOGIN_PATH              = "$"
	// database message success & error
	DATABASE_GET_SUCCESS    = "database search successful"
	DATABASE_GET_FAIL       = "database search failed"
	DATABASE_APPEND_SUCCESS = "database new data added"
	DATABASE_APPEND_FAIL    = "fail to append new data"
	DATABASE_CREATE_SUCCESS = "database new key-value pair added"
	DATABASE_CREATE_FAIL    = "fail to create new key-value pair"
	DATABASE_DELETE_SUCCESS = "database requested data deleted"
	DATABASE_DELETE_FAIL    = "fail to delete requested data"
	// request messages
	REQUEST_WRONG_TYPE    = "wrong type parameter"
	REQUEST_MISSING_PARAM = "missing parameter"
	// price information
	PRICE_GET_ERROR = "binance failure"
	// login information
	LOGIN_SUCCESS = "login successful"
	LOGIN_ERROR   = "login fail"
)

var EmployTypeMap = map[string]string{
	"perma_key":  PERMANENT_EMPLOYER_KEY,
	"perma_path": PERMANENT_EMPLOYER_PATH,
	"free_key":   FREELANCE_EMPLOYER_KEY,
	"free_path":  FREELANCE_EMPLOYER_PATH,
}
