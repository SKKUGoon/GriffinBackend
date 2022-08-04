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