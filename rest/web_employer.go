package rest

import (
	"GriffinBackend/rdb"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func postEmployer(c *gin.Context, db *redis.Client) {
	newEmployer, err := handleParamEmployerId(c)
	if err != nil {
		return
	}
	ph := rdb.Employee{
		Id:   0,
		Name: "initial placeholder",
	}
	newInit := []rdb.Employee{ph}
	_ = rdb.JsonSet(
		db,
		fmt.Sprintf(PERMANENT_EMPLOYER_KEY, newEmployer),
		PERMANENT_EMPLOYER_PATH,
		&newInit,
	)
	_ = rdb.JsonSet(
		db,
		fmt.Sprintf(FREELANCE_EMPLOYER_KEY, newEmployer),
		FREELANCE_EMPLOYER_PATH,
		&newInit,
	)
	c.JSON(http.StatusOK, gin.H{
		"message": DATABASE_CREATE_SUCCESS,
	})
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
