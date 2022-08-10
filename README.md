<div align="center">

<img src="https://user-images.githubusercontent.com/54009920/183289711-bf7c3303-0cf6-48ee-a270-d777d8d1bc45.svg" alt="icon" width="250"/>

<br/>
📆 2022.07.31 ~ Ongoing

</div>

<br/>
<br/>

<div align="center">

[![GitHub Open Issues](https://img.shields.io/github/issues-raw/skkugoon/GriffinBackend?color=green)](https://github.com/skkugoon/GriffinBackend/issues)
[![GitHub Closed Issues](https://img.shields.io/github/issues-closed-raw/skkugoon/GriffinBackend?color=red)](https://github.com/skkugoon/GriffinBackend/issues?q=is%3Aissue+is%3Aclosed)
[![GitHub Open PR](https://img.shields.io/github/issues-pr-raw/skkugoon/GriffinBackend?color=green)](https://github.com/skkugoon/GriffinBackend/pulls)
[![GitHub Closed PR](https://img.shields.io/github/issues-pr-closed-raw/skkugoon/GriffinBackend?color=red)](https://github.com/skkugoon/GriffinBackend/pulls?q=is%3Apr+is%3Aclosed)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
  
</div>

# Griffin Backend

## File Tree
```bash
# file tree
├── price
│   ├── binance_info.go
│   └── constant.go
├── rdb
│   ├── dao_conn.go
│   ├── dao_struct.go
│   └── re_json.go
├── rdb
│   ├── constant.go
│   ├── src_serve.go
│   ├── src_util.go
│   ├── web_currency.go
│   ├── web_employee.go
│   ├── web_employer.go
│   ├── web_handler.go
│   ├── web_landing.go
│   └── web_param.go
├── .env
├── .gitignore
├── CHANGELOG.md
├── Dockerfile
├── go.mod
├── LICENSE
├── README.go
└── main.go
```

## Requirements
- go 1.18+
- docker 4.11.0+

## Installation
```bash
$ go mod init GriffinBackend
$ go mod download
```
## Griffin Backend

## Endpoint Tree
```bash
# API endpoints
├── /ping
│   └── GET
├── /version
│   └── GET
├── /login
│   └── GET
├── /payment
│   ├── GET
│   └── POST
├── /paymentMonth
│   └── GET
├── /employer
│   ├── GET
│   └── POST
├── /employee
│   ├── GET
│   ├── POST
│   └── DELETE
└── /price
    └── GET

```
## API Endpoint Explanation
## Ping
| Type | Value |
|------|------|
| Endpoint | `/ping` |
| Method | `GET` |

```bash
$ curl http://localhost:8000/ping
```
```json
// Return
{
	"message": "pong"
} 
```
<p>
API Health check. Returns JSON with message.
</p>

## Version
| Type | Value |
|------|------|
| Endpoint | `/version` |
| Method | `GET` |

```bash
$ curl http://localhost:8000/version
```

```json
// Return
{
	"message": "0.0.5"
}
```
<p>
API Version. Gets it from .env file
</p>

## Login
| Type | Value |
|------|------|
| Endpoint | `/login` |
| Method | `GET` |

```bash
$ curl http://localhost:8000/login?employerId=5&employerPw=abcde
```
```json
// Return
{
	"employerId": "5",
	"employerPw": "abcde",
	"message": "login successful"
}
```

## GetPayment
| Type | Value |
|------|------|
| Endpoint | `/payment` |
| Method | `GET` |

| Queries | Example (type) |
|------|------|
| employerId | 5 (int) |
```bash
$ curl http://localhost:8000/payment?employerId=5
```
```json
// Return
[
	{
		"name": "pia",
		"payroll": 100000,
		"curr": "usdc",
		"time": "2022-08-07T01:36:02+0900"
	},
	{
		"name": "pia",
		"payroll": 12,
		"curr": "usdc",
		"time": "2022-08-07T01:36:07+0900"
	},
	{
		"name": "pia",
		"payroll": 12000000,
		"curr": "matic",
		"time": "2022-08-07T05:26:06+0900"
	}
]
```
<p>

Get payment record for a single employer identified by `employerId`

</p>

## PostPayment
| Type | Value |
|------|------|
| Endpoint | `/payment` |
| Method | `POST` |

| Queries | Example (type) |
|------|------|
| employerId | 5 (int) |
| name | pia (str) |
| payroll | 12000000 (float64) |
| curr | matic (str) |
```bash
$ curl http://localhost:8000/payment?employerId=5&name=pia&payroll=12000000&curr=matic
```
```json
// Return
{
	"message": "database new data added"
}
```
## GetPaymentMonth
| Type | Value |
|------|------|
| Endpoint | `/paymentMonth` |
| Method | `GET` |

| Queries | Example (type) |
|------|------|
| employerId | 5 (int) |
```bash
$ curl http://localhost:8000/paymentMonth?employerId=5
```
```json
// Return
[
	{
		"date": "2022-08",
		"eth": 0,
		"matic": 12000000,
		"usdc": 100012
	}
]
```

<p>

Get monthly aggregated payment information

</p>

## PostEmployer
| Type | Value |
|------|------|
| Endpoint | `/employer` |
| Method | `POST` |

| Queries | Example (type) |
|------|------|
| employerId | 5 (int) |
| employerPw | abcde (str) |

```bash
$ curl http://localhost:8000/employer?employerId=5&employerPw=abcde
```
```json
// Return
{
	"employerId": "5",
	"employerPw": "abcde",
	"message": "database new key-value pair added"
}
```
<p>
POST employer would automatically create 3 key-value json tables inside redis database;

  - employee_perma:'employeeId'
  - employee_free:'employeeId'
  - payroll_history:'employeeId'
</p>

## DeleteEmployer
| Type | Value |
|----|----|
| Endpoint | `employer` |
| Method | `DELETE` |

| Queries | Example (type) |
| --- | --- |
| employerId | 1 (int) |

```bash
$ curl http://localhost:8000/employer?employerId=1
```

```json
// Return
{
	"message": "database requested data deleted"
}
```

<p>

Delete employer database. Employer identified by `employerId`

</p>

## GetEmployee
| Type | Value |
|------|------|
| Endpoint | `/ping` |
| Method | `GET` |

| Queries | Example (type) |
|------|------|
| employerId | 5 (int) |
```bash
$ curl http://localhost:8000/employee?employerId=5
```
```json
// Return
{
	"total_length": 7,
	"repeat_true": [
		{
			"key": 2203,
			"name": "julia bae",
			"email": "juljul123@jul.xyz",
			"position": "smart contract dev",
			"account": "0x01you",
			"payroll": 10000,
			"curr": "usdc",
			"date": 23,
			"secLeft": 1088987.461009
		},
		{
			"key": 2205,
			"name": "julia bae",
			"email": "juljul123@jul.xyz",
			"position": "smart contract dev",
			"account": "0x01you",
			"payroll": 10000,
			"curr": "usdc",
			"date": 23,
			"secLeft": 1088987.460967
		},
		{
			"key": 2203,
			"name": "julia bae",
			"email": "juljul123@jul.xyz",
			"position": "smart contract dev",
			"account": "0x01you",
			"payroll": 10000,
			"curr": "usdc",
			"date": 23,
			"secLeft": 1088987.460962
		},
		{
			"key": 2208,
			"name": "julia bae",
			"email": "juljul123@jul.xyz",
			"position": "smart contract dev",
			"account": "0x01you",
			"payroll": 10000,
			"curr": "usdc",
			"date": 23,
			"secLeft": 1088987.460959
		}
	],
	"repeat_false": [
		{
			"key": 2108,
			"name": "julia bae",
			"email": "juljul123@jul.xyz",
			"position": "smart contract dev",
			"account": "0x01you",
			"payroll": 10000,
			"curr": "usdc",
			"date": 23,
			"secLeft": 1088987.460956
		},
		{
			"key": 2103,
			"name": "julia bae",
			"email": "juljul123@jul.xyz",
			"position": "smart contract dev",
			"account": "0x01you",
			"payroll": 10000,
			"curr": "usdc",
			"date": 23,
			"secLeft": 1088987.460952
		},
		{
			"key": 2122,
			"name": "julia bae",
			"email": "juljul123@jul.xyz",
			"position": "smart contract dev",
			"account": "0x01you",
			"payroll": 10000,
			"curr": "usdc",
			"date": 23,
			"secLeft": 1088987.460949
		}
	]
}
```
<p>

Get all the employees for employer identified by their employerId `employerId`
</p>

## PostEmployee
| Type | Value |
|------|------|
| Endpoint | `/employee` |
| Method | `POST` |

| Queries | Example (type) |
|------|------|
| employerId | 1 (int) |
| employType | free (string) |
| key | 2122 (int) |
| name | Julia Bae (str) |
| email | juljul123@jul.xyz (str) |
| position | smart contract dev (str)  |
| account | 0x01you (str) |
| curr | usdc (str)  |
| payroll | 10000 (float64) |
| date | 23 (int) |
```bash
curl http://localhost:8000/employee?employerId=5&employType=free&key=2122&name=julia%20bae&email=juljul123%40jul.xyz&position=smart%20contract%20dev&account=0x01you&curr=usdc&payroll=10000&date=23
```
```json
// Return
{
	"message": "database new data added"
}
```
<p>

Add new employee with unique information regarding their payments etc.

</p>

## DeleteEmployee
| Type | Value |
|------|------|
| Endpoint | `/employee` |
| Method | `DELETE` |

| Queries | Example (type) |
|------|------|
| employerId | 1 (int) |
| id | 5 (int) |
| employType | free (str) |
```bash
$ curl http://localhost:8000/employee?employerId=1&id=5&employType=free
```
```json
// Return
Null
```
## GetPrice
| Type | Value |
|------|------|
| Endpoint | `/price` |
| Method | `GET` |

```bash
curl http://localhost:8000/price
```
```json
// Return
{
	"ethusdt": 1714.02,
	"maticusdt": 0.9204,
	"usdcusdt": 0.9999
}
```

<p>

Get currency information from Binance API.

- ETHUSDT
- MATICUSDT
- USDCUSDT

</p>
