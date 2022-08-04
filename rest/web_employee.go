package rest

import (
	"GriffinBackend/rdb"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func getEmployee(c *gin.Context, db *redis.Client) {
	var (
		employeeP [][]rdb.Employee
		employeeF [][]rdb.Employee

		p []rdb.Employee
		f []rdb.Employee
	)

	eid, err := handleParamEmployerId(c)
	if err != nil {
		return
	}
	// permenent work force
	e0 := rdb.JsonGet(db, fmt.Sprintf(PERMANENT_EMPLOYER_KEY, eid), PERMANENT_EMPLOYER_PATH, &employeeP)
	// temporary work force
	e1 := rdb.JsonGet(db, fmt.Sprintf(FREELANCE_EMPLOYER_KEY, eid), FREELANCE_EMPLOYER_PATH, &employeeF)
	switch {
	// [1:] because first content is always a placeholder
	case e0 != nil:
		p, f = []rdb.Employee{}, employeeF[0][1:]
	case e1 != nil:
		p, f = employeeP[0][1:], []rdb.Employee{}
	case (e0 != nil) && (e1 != nil):
		p, f = []rdb.Employee{}, []rdb.Employee{}
	default:
		p, f = employeeP[0][1:], employeeF[0][1:]
	}
	// send it to webserver
	totalEmp := rdb.EmployeePlural{
		Total:     len(employeeF[0]) + len(employeeP[0]) - 2,
		Permanent: p,
		Freelance: f,
	}
	c.JSON(http.StatusOK, totalEmp)
}

/**
postEmployee
- required query
  - employerId
  - employType

  - name
  - email
  - position
  - account
  - curr
  - payroll
  - date
*/
func postEmployee(c *gin.Context, db *redis.Client) {
	newEmployee, err := handleNewEmployeeInfo(c)
	if err != nil {
		return
	}
	onboardInfo, err := handleNewEmployeeID(c, db)
	kK := fmt.Sprintf("%v_%v", onboardInfo.EmploymentType, "key")
	kP := fmt.Sprintf("%v_%v", onboardInfo.EmploymentType, "path")

	newEmployee.Id = onboardInfo.MyEmployeeId
	err = rdb.JsonArrAppend(
		db,
		fmt.Sprintf(EmployTypeMap[kK], onboardInfo.MyEmployerId),
		EmployTypeMap[kP],
		&newEmployee,
	)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"message": DATABASE_APPEND_FAIL,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": DATABASE_APPEND_SUCCESS,
		})
	}
}

func deleteEmployee(c *gin.Context, db *redis.Client) {
	handleDelEmployee(c, db)
	c.JSON(http.StatusOK, gin.H{
		"message": DATABASE_DELETE_SUCCESS,
	})
}

func updateEmployee(c *gin.Context, db rdb.DataBase) {

}
