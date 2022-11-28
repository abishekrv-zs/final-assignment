package company

import (
	"database/sql"
	"github.com/abishekrv-zs/final-assignment/model"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return Store{db: db}
}

func (s Store) GetAll() ([]model.Company, error) {
	return nil, nil
}

func (s Store) GetById(id string) (model.Company, error) {
	return model.Company{}, nil
}

func (s Store) Create(c model.Company) (model.Company, error) {
	return model.Company{}, nil
}

func (s Store) Update(id string) (model.Company, error) {
	return model.Company{}, nil
}

func (s Store) Delete(id string) (model.Company, error) {
	return model.Company{}, nil
}
