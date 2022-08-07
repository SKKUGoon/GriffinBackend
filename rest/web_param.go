package rest

import (
	"GriffinBackend/rdb"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func handleParamEmployerId(c *gin.Context) (string, error) {
	// check for employer id
	q, ok := c.GetQuery(EMPLOYER_ID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": REQUEST_MISSING_PARAM + " " + EMPLOYER_ID,
		})
		return "", errors.New(REQUEST_MISSING_PARAM)
	}
	return q, nil
}

func handleParamEmployerPw(c *gin.Context) (string, error) {
	// check for employer password
	p, ok := c.GetQuery(EMPLOYER_PW)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": REQUEST_MISSING_PARAM + " " + EMPLOYER_PW,
		})
		return "", errors.New(REQUEST_MISSING_PARAM)
	}
	return p, nil
}

func handleParamEmployType(c *gin.Context) (string, error) {
	// check for employment type
	q, ok := c.GetQuery(EMPLOYMENT_TYPE)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": REQUEST_MISSING_PARAM + " " + EMPLOYMENT_TYPE,
		})
		return "", errors.New(REQUEST_MISSING_PARAM)
	}
	return q, nil
}

func handleParamEmployeePartial(c *gin.Context) bool {
	_, ok := c.GetQuery(EMPLOYEE_PARTIAL)
	if !ok {
		return false
	}
	return true
}

func handleParamDelEmployee(c *gin.Context) (int, error) {
	q, ok := c.GetQuery(EMPLOYEE_ID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": REQUEST_MISSING_PARAM + " " + EMPLOYEE_ID,
		})
		return -1, errors.New(REQUEST_MISSING_PARAM)
	}
	qInt, err := strconv.Atoi(q)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": REQUEST_WRONG_TYPE + " " + EMPLOYEE_ID,
		})
	}
	return qInt, nil
}

func handleParamPostPay(c *gin.Context) (rdb.Payment, error) {
	name, ok0 := c.GetQuery(EMPLOYEE_NAME)
	payroll, ok1 := c.GetQuery(EMPLOYEE_PAYROLL)
	currency, ok2 := c.GetQuery(EMPLOYEE_CURRENCY)

	if !ok0 || !ok1 || !ok2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": REQUEST_MISSING_PARAM,
		})
		return rdb.Payment{}, errors.New(REQUEST_MISSING_PARAM)
	}

	payrollInt, _ := strconv.Atoi(payroll)
	t := time.Now().Format("2006-01-02T15:04:05-0700")
	p := rdb.Payment{
		Name:     name,
		Payroll:  payrollInt,
		Currency: currency,
		Time:     t,
	}
	return p, nil
}
