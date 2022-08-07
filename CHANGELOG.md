# 0.0.1

<p>

First Commit. Create Empty files + directory

</p>

[Add]
- ./rdb
  - dao_conn.go
  - dao_struct.go
  - re_json.go
- ./rest
  - constant.go
  - src_serve.go
- other necessary files such as .gitignore, CHANGELOG etc.

[Change]


[Fix]


[Remove]


# 0.0.2

<p>

Add JSON database structure. TODO: redis
</p>

[Add]
- ./rdb
  - dao_conn.go
    - struct DataBase
      - method Connect 
  - dao_struct.go
    - struct Employee
    - struct EmployeePlural
- ./main.go
  - func init - load environment
  
[Change]
- ./Dockerfile

[Fix]

[Remove]

# 0.0.3
<p>

Complete JSON database structure.
</p>

[Add]
- ./rdb
  - dao_conn.go
    - interface DAO
    - structure DataBase
  - re_json.go
    - func JsonGet
    - func JsonSet
    - func JsonArrAppend
- ./rest
  - web_employ.go
    - func getEmployee - get from redis database
    - func postEmployee - insert into redis database
    - func deleteEmployee - delete from redis database
    - func updateEmployee - update redis database information 
  - web_landing.go
    - func pingPong - pingPong connection test
    - func version - gives api version
  - src_serve.go
    - basic methods
      - method StartService  - /  GET
      - method PingTest  - /ping  GET
      - method Version  - /version  GET
    - employ methods
      - method GetEmployee  - /employee  GET 
      - method AddEmployee  - /employee  POST

[Change]


[Fix]
- ./rdb
  - dao_conn.go
    - func Connect -> returning *redis.Client. Fixed from DataBase Method

[Remove]


# 0.0.4
<p>

cr employer, crud employee and docker
</p>

[Add]
- ./main.go
  - add AddEmployer() and DeleteEmployer()
  - add AddEmployee() and DeleteEmployee()
- ./rdb
  - dao_struct.go
    - struct EmployeeOnboard - new employee information
  - re_json.go
    - func JsonDel
- ./rest
  - constant.go
    - Add constant - EMPLOYEE_CURRENCY, EMPLOYEE_CURRENCY_INDEX and messages
  - src_serve.go
    - method AddEmployer
    - method DeleteEmployer
    - method DeleteEmployee
  - web_employee.go
    - func getEmployee
    - func postEmployee
    - func deleteEmployee
  - web_param.go
    - func handleParamEmployerId
    - func handleParamEmployType
    - func handleParamDelEmployee
  - web_handler.go
    - func handleNewEmployeeID
    - func handleNewEmployeeInfo
    - func handleDelEmployee
  - web_employer.go
    - func postEmployer
    - func delEmployer


[Change]

[Fix]

[Remove]

# 0.0.5

`7cbe591`

<p>
Fix cors problem
</p>

[Add]
- ./rest
  - constant.go
    - access control constants - allow origin, credentials
  - src.utils.go
    - func CORSMiddleware

[Change]
- main.go
  - griffinPay additional settings - Readtimeout, WriteTimeout etc.

[Fix]
- ./rest
  - src_serve.go
    - method StartService - gin connect uses CORS Middleware at web service start
    
[Remove]

<p>
Get price information from binance
</p>

[Add]
- ./price
  - binance_info.go
    - func BinancePrice - main function
    - func buildRequest[T any] - insert queries (inside a sturct) into a url 
    - func structIter - utility function that iterates through a string
  - constant.go
    - struct BinanceRequest
    - struct BinanceResponse
    - struct PriceInformation
    - and other constant
  
- ./rest 
  - src_serve.go
    - func GetPrice
  - web.currency.go
    - func getBinanceTrade
```json
{
	"ethusdt": 1718.67,
	"maticusdt": 0.9317,
	"usdcusdt": 0.9999
}
```

[Change]

[Fix]

[Remove]


`da0e677`

<p>

recently added worker | historical payroll activity | login page

</p>

[Add]
- ./rdb
  - dao_struct.go
    - struct Payment - name payroll currency and time
- ./rest
  - constant.go
    - EMPLOYEE_PARTIAL
    - HISTORICAL_PAYMENT_KEY
    - HISTORICAL_PAYMENT_PATH
    - LOGIN_KEY
    - LOGIN_PATH
    - DATABASE_GET_SUCCESS
    - DATABASE_GET_FAIL
  - src_serve.go
    - method AddPaymentRecord
    - method GetPaymentRecord
  - src_util.go
    - func processEmployee
    - func calcTimeLeft
  - web_param.go
    - func handleParamEmployerPw
    - func handleParamEmployeePartial   - partial employ list true or false
    - func handleParamPostPay  - new payment history
  

[Change]
- ./price
  - binance_info.go
    - func BinancePrice - delete fmt.Print
    - func buildRequest[T any] - delete fmt.Print
- ./rest
  - constant.go
    - EMPLOYEE_ID is now key
    - EMPLOYEE_PARTIAL
  - web_employee.go 
    - func getEmployee
      - if handleParamEmployeePartial is true - give recent employees (2, 2, total of 4)
      - processEmployee(e) -> calculate seconds left till payment
  - web_employer.go
    - func postEmployeer
      - add password when inserting employer
      - add new table - historical payment recording purpose
  - web_currency.go
    - func postPayment
    - func getPayment


`28119d0`

<p>

add login method
</p>

[Add]
- ./rest
  - web_employer.go
    - func loginEmployer - StatusUnauthorized if wrong password | StatusForbidden if wrong id
  - src_util.go
    - func structIter - same as one in price / binance_info.go
    - func isRegistered - true true if all success
  - src_serve.go
    - method Login
  - constants.go
    - LOGIN_SUCCESS, LOGIN_ERROR constant

- ./rdb
  - dao_struct.go
    - struct Login  

[Change]
- ./rest
  - web_employer.go
    - func postEmployer - newIdPw and add section to id and pw

[Fix]

[Remove]

`a104c47`

<p>
Stop working
</p>

[Add]
- ./price
  - constant.go
    - struct MonthlyPayHistory
- ./rest
  - web_currency.go
    - func getPaymentMonthly
  - src_util.go
    - func monthlyPayment
    - func monthlyPaymentStruct
  - src_serve.go
    - func GetPaymentRecordMonth

[Change]

[Fix]

[Remove]