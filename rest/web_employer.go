package rest

import (
	"GriffinBackend/rdb"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func postEmployer(c *gin.Context, db *redis.Client) {
	// queries for new employer
	newEmployer, err := handleParamEmployerId(c)
	if err != nil {
		return
	}
	newPassword, err := handleParamEmployerPw(c)
	if err != nil {
		return
	}

	phEmp := rdb.Employee{
		Id:   0,
		Name: "initial placeholder",
	}
	phPay := rdb.Payment{}

	newInitEmp := []rdb.Employee{phEmp}
	newInitPay := []rdb.Payment{phPay}
	newIdPw := rdb.Login{
		Username: newEmployer,
		Password: newPassword,
	}
	// permenent employee data table
	_ = rdb.JsonSet(
		db,
		fmt.Sprintf(PERMANENT_EMPLOYER_KEY, newEmployer),
		PERMANENT_EMPLOYER_PATH,
		&newInitEmp,
	)
	// freelance employee data table
	_ = rdb.JsonSet(
		db,
		fmt.Sprintf(FREELANCE_EMPLOYER_KEY, newEmployer),
		FREELANCE_EMPLOYER_PATH,
		&newInitEmp,
	)
	// employee historical payment data table
	_ = rdb.JsonSet(
		db,
		fmt.Sprintf(HISTORICAL_PAYMENT_KEY, newEmployer),
		HISTORICAL_PAYMENT_PATH,
		&newInitPay,
	)
	// add employee id pw to login table
	_ = rdb.JsonArrAppend(
		db,
		LOGIN_KEY,
		LOGIN_PATH,
		&newIdPw,
	)
	c.JSON(http.StatusOK, gin.H{
		"message":    DATABASE_CREATE_SUCCESS,
		"employerId": newEmployer,
		"employerPw": newPassword,
	})
}

func loginEmployer(c *gin.Context, db *redis.Client) {
	username, err := handleParamEmployerId(c)
	if err != nil {
		return
	}
	password, err := handleParamEmployerPw(c)
	if err != nil {
		return
	}

	var registered [][]rdb.Login
	_ = rdb.JsonGet(
		db,
		LOGIN_KEY,
		LOGIN_PATH,
		&registered,
	)
	id, pw := isRegistered(registered[0], username, password)
	switch {
	case id && pw:
		c.JSON(http.StatusOK, gin.H{
			"message":    LOGIN_SUCCESS,
			"employerId": username,
			"employerPw": password,
		})
	case id:
		// if wrong password StatusUnauthorized
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": LOGIN_ERROR + " wrong password",
		})
	default:
		// if no id StatusForbidden
		c.JSON(http.StatusForbidden, gin.H{
			"message": LOGIN_ERROR + " id unrecognized",
		})
	}
}

func delEmployer(c *gin.Context, db *redis.Client) {
	Employer, err := handleParamEmployerId(c)
	if err != nil {
		return
	}
	_ = rdb.JsonDel(
		db,
		fmt.Sprintf(PERMANENT_EMPLOYER_KEY, Employer),
		PERMANENT_EMPLOYER_PATH,
	)
	_ = rdb.JsonDel(
		db,
		fmt.Sprintf(FREELANCE_EMPLOYER_KEY, Employer),
		FREELANCE_EMPLOYER_PATH,
	)
	c.JSON(http.StatusOK, gin.H{
		"message": DATABASE_DELETE_SUCCESS,
	})
}
