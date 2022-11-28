package svc

import "github.com/abishekrv-zs/final-assignment/model"

type CompanySvc interface {
	getAll() ([]model.Company, error)
	getById(id string) (model.Company, error)
	create(c model.Company) (model.Company, error)
	update(id string) (model.Company, error)
	delete(id string) (model.Company, error)
}

type StudentSvc interface {
	getAll() ([]model.Student, error)
	getById(id string) (model.Student, error)
	create(s model.Student) (model.Student, error)
	update(id string) (model.Student, error)
	delete(id string) (model.Student, error)
}
