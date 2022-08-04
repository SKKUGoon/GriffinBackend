package rest

import (
	"errors"
	"net/http"
	"strconv"

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
