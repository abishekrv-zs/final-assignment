package company

import (
	"database/sql"
	"github.com/abishekrv-zs/final-assignment/model"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) store {
	return store{db: db}
}

func (s store) GetAll() ([]model.Company, error) {
	return nil, nil
}

func (s store) GetById(id string) (model.Company, error) {
	return model.Company{}, nil
}

func (s store) Create(c model.Company) (model.Company, error) {
	return model.Company{}, nil
}

func (s store) Update(id string) (model.Company, error) {
	return model.Company{}, nil
}

func (s store) Delete(id string) (model.Company, error) {
	return model.Company{}, nil
}
