package student

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

func (s Store) GetAll() ([]model.Student, error) {
	return nil, nil
}

func (s Store) GetById(id string) (model.Student, error) {
	return model.Student{}, nil
}

func (s Store) Create(student model.Student) (model.Student, error) {
	return model.Student{}, nil
}

func (s Store) Update(id string) (model.Student, error) {
	return model.Student{}, nil
}

func (s Store) Delete(id string) (model.Student, error) {
	return model.Student{}, nil
}
