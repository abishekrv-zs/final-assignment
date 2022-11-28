package student

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

func (s store) GetAll() ([]model.Student, error) {
	return nil, nil
}

func (s store) GetById(id string) (model.Student, error) {
	return model.Student{}, nil
}

func (s store) Create(student model.Student) (model.Student, error) {
	return model.Student{}, nil
}

func (s store) Update(id string) (model.Student, error) {
	return model.Student{}, nil
}

func (s store) Delete(id string) (model.Student, error) {
	return model.Student{}, nil
}
