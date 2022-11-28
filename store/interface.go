package store

import "github.com/abishekrv-zs/final-assignment/model"

type CompanyStore interface {
	GetAll() ([]model.Company, error)
	GetById(id string) (model.Company, error)
	Create(c model.Company) (model.Company, error)
	Update(id string) (model.Company, error)
	Delete(id string) (model.Company, error)
}

type StudentStore interface {
	GetAll() ([]model.Student, error)
	GetById(id string) (model.Student, error)
	Create(s model.Student) (model.Student, error)
	Update(id string) (model.Student, error)
	Delete(id string) (model.Student, error)
}
