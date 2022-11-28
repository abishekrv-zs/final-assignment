package main

import (
	"database/sql"
	deliveryStudent "github.com/abishekrv-zs/final-assignment/delivery/student"
	storeStudent "github.com/abishekrv-zs/final-assignment/store/student"
	svcStudent "github.com/abishekrv-zs/final-assignment/svc/student"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "connectionStirng")
	if err != nil {
		log.Println(err)
		return
	}

	store := storeStudent.New(db)
	svc := svcStudent.New(store)
	h := deliveryStudent.New(svc)

	router := mux.NewRouter()
	router.HandleFunc("/students", h.GetAll).Methods("GET")
	router.HandleFunc("/students/{id}", h.GetById).Methods("GET")
	router.HandleFunc("/students", h.Create).Methods("POST")
	router.HandleFunc("/students", h.Update).Methods("PUT")
	router.HandleFunc("/students", h.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}
