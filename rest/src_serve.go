package rest

import (
	"GriffinBackend/rdb"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type GriffinWS struct {
	Conn     *gin.Engine
	Database *redis.Client
}

func (g GriffinWS) StartService() GriffinWS {
	// start web server connection
	g.Conn = gin.Default()
	// add database initialization
	d := rdb.Connect()
	g.Database = d
	return g
}

func (g GriffinWS) PingTest() GriffinWS {
	g.Conn.GET("/ping", pingPong)
	return g
}

func (g GriffinWS) Version() GriffinWS {
	g.Conn.GET("/version", version)
	return g
}

func (g GriffinWS) AddEmployer() GriffinWS {
	g.Conn.POST("/employer", func(c *gin.Context) {
		postEmployer(c, g.Database)
	})
	return g
}

func (g GriffinWS) DeleteEmployer() GriffinWS {
	g.Conn.DELETE("/employer", func(c *gin.Context) {
		delEmployer(c, g.Database)
	})
	return g
}

func (g GriffinWS) AddEmployee() GriffinWS {
	g.Conn.POST("/employee", func(c *gin.Context) {
		postEmployee(c, g.Database)
	})
	return g
}

func (g GriffinWS) GetEmployee() GriffinWS {
	g.Conn.GET("/employee", func(c *gin.Context) {
		getEmployee(c, g.Database)
	})
	return g
}

func (g GriffinWS) DeleteEmployee() GriffinWS {
	g.Conn.DELETE("/employee", func(c *gin.Context) {
		deleteEmployee(c, g.Database)
	})
	return g
}
