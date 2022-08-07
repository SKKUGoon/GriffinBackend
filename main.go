package main

import (
	"GriffinBackend/rest"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}
}

var griffin = rest.GriffinWS{}

func main() {
	// landing page
	griffin = griffin.
		StartService().
		PingTest().
		Version().
		Login()

	// employer operations - preferably admin's job
	griffin = griffin.
		AddEmployer().
		DeleteEmployer()

	// employee CRUD operations
	griffin = griffin.
		GetEmployee().
		AddEmployee().
		DeleteEmployee()

	// currency get operations
	griffin = griffin.
		GetPrice().
		AddPaymentRecord().
		GetPaymentRecord().
		GetPaymentRecordMonth()

	griffinPay := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        griffin.Conn,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	griffinPay.ListenAndServe()
}
