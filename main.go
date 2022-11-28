package main

import (
	"database/sql"
	deliveryCompany "github.com/abishekrv-zs/placement-api/delivery/company"
	deliveryStudent "github.com/abishekrv-zs/placement-api/delivery/student"
	storeCompany "github.com/abishekrv-zs/placement-api/store/company"
	storeStudent "github.com/abishekrv-zs/placement-api/store/student"
	svcCompany "github.com/abishekrv-zs/placement-api/svc/company"
	svcStudent "github.com/abishekrv-zs/placement-api/svc/student"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func setDbConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "connectionStirng")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

func main() {

	db, err := setDbConnection()
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	studentStore := storeStudent.New(db)
	studentSvc := svcStudent.New(studentStore)
	studentHandler := deliveryStudent.New(studentSvc)

	companyStore := storeCompany.New(db)
	companySvc := svcCompany.New(companyStore)
	companyHandler := deliveryCompany.New(companySvc)

	router := mux.NewRouter()
	router.HandleFunc("/students", studentHandler.GetAll).Methods("GET")
	router.HandleFunc("/students/{id}", studentHandler.GetById).Methods("GET")
	router.HandleFunc("/students", studentHandler.Create).Methods("POST")
	router.HandleFunc("/students", studentHandler.Update).Methods("PUT")
	router.HandleFunc("/students", studentHandler.Delete).Methods("DELETE")

	router.HandleFunc("/companies", companyHandler.GetAll).Methods("GET")
	router.HandleFunc("/companies/{id}", companyHandler.GetById).Methods("GET")
	router.HandleFunc("/companies", companyHandler.Create).Methods("POST")
	router.HandleFunc("/companies", companyHandler.Update).Methods("PUT")
	router.HandleFunc("/companies", companyHandler.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}
