package company

import (
	"github.com/abishekrv-zs/final-assignment/svc"
	"net/http"
)

type Delivery struct {
	svc svc.CompanySvc
}

func New(svc svc.CompanySvc) Delivery {
	return Delivery{svc: svc}
}

func (h Delivery) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (h Delivery) GetById(w http.ResponseWriter, r *http.Request) {

}

func (h Delivery) Create(w http.ResponseWriter, r *http.Request) {

}

func (h Delivery) Update(w http.ResponseWriter, r *http.Request) {

}

func (h Delivery) Delete(w http.ResponseWriter, r *http.Request) {

}
