package rest

import (
	"GriffinBackend/rdb"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func handleNewEmployeeID(c *gin.Context, db *redis.Client) (rdb.EmployeeOnboard, error) {
	eid, err := handleParamEmployerId(c)
	if err != nil {
		return rdb.EmployeeOnboard{}, err
	}
	wid, err := handleParamDelEmployee(c)
	if err != nil {
		return rdb.EmployeeOnboard{}, err
	}
	etype, err := handleParamEmployType(c)
	if err != nil {
		return rdb.EmployeeOnboard{}, err
	}

	if err != nil {
		return rdb.EmployeeOnboard{}, err
	}

	result := rdb.EmployeeOnboard{
		MyEmployerId:   eid,
		MyEmployeeId:   wid,
		EmploymentType: etype,
	}
	return result, nil
}

func handleNewEmployeeInfo(c *gin.Context) (rdb.Employee, error) {
	needQuery := []string{
		EMPLOYEE_NAME, EMPLOYEE_EMAIL, EMPLOYEE_POSITION, EMPLOYEE_ACCOUNT,
		EMPLOYEE_PAYROLL, EMPLOYEE_PAYDAY, EMPLOYEE_CURRENCY,
	}
	needValue := []string{}

	for _, qs := range needQuery {
		q, ok := c.GetQuery(qs)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": REQUEST_MISSING_PARAM + " " + qs,
			})
			return rdb.Employee{}, errors.New(REQUEST_MISSING_PARAM)
		}
		needValue = append(needValue, q)
	}

	payroll, err := strconv.Atoi(needValue[EMPLOYEE_PAYROLL_INDEX])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messsage": REQUEST_WRONG_TYPE + " payroll",
		})
		return rdb.Employee{}, errors.New(REQUEST_WRONG_TYPE)
	}
	date, err := strconv.Atoi(needValue[EMPLOYEE_PAYDAY_INDEX])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": REQUEST_WRONG_TYPE + " date",
		})
		return rdb.Employee{}, errors.New(REQUEST_WRONG_TYPE)
	}

	return rdb.Employee{
		Name:     needValue[EMPLOYEE_NAME_INDEX],
		Email:    needValue[EMPLOYEE_EMAIL_INDEX],
		Position: needValue[EMPLOYEE_POSITION_INDEX],
		Account:  needValue[EMPLOYEE_ACCOUNT_INDEX],
		Payroll:  payroll,
		Currency: needValue[EMPLOYEE_CURRENCY_INDEX],
		Date:     date,
	}, nil
}

func handleDelEmployee(c *gin.Context, db *redis.Client) {
	newSet := []rdb.Employee{}

	eid, err := handleParamEmployerId(c)
	if err != nil {
		return
	}
	wid, err := handleParamDelEmployee(c)
	if err != nil {
		return
	}
	etype, err := handleParamEmployType(c)
	if err != nil {
		return
	}
	kK := fmt.Sprintf("%v_%v", etype, "key")
	kP := fmt.Sprintf("%v_%v", etype, "path")
	// get all the data
	var result [][]rdb.Employee
	err = rdb.JsonGet(
		db,
		fmt.Sprintf(EmployTypeMap[kK], eid),
		EmployTypeMap[kP],
		&result,
	)

	fmt.Println(result)
	for _, i := range result[0] {
		if i.Id != wid {
			newSet = append(newSet, i)
		}
	}
	err = rdb.JsonSet(
		db,
		fmt.Sprintf(EmployTypeMap[kK], eid),
		EmployTypeMap[kP],
		&newSet,
	)
}
