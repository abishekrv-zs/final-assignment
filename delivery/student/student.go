package student

import (
	"github.com/abishekrv-zs/final-assignment/svc"
	"net/http"
)

type delivery struct {
	svc svc.StudentSvc
}

func New(svc svc.StudentSvc) delivery {
	return delivery{svc: svc}
}

func (h delivery) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (h delivery) GetById(w http.ResponseWriter, r *http.Request) {

}

func (h delivery) Create(w http.ResponseWriter, r *http.Request) {

}

func (h delivery) Update(w http.ResponseWriter, r *http.Request) {

}

func (h delivery) Delete(w http.ResponseWriter, r *http.Request) {

}
