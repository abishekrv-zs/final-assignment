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

func (s store) getAll() ([]model.Student, error) {
	return nil, nil
}

func (s store) getById(id string) (model.Student, error) {
	return model.Student{}, nil
}

func (s store) create(student model.Student) (model.Student, error) {
	return model.Student{}, nil
}

func (s store) update(id string) (model.Student, error) {
	return model.Student{}, nil
}

func (s store) delete(id string) (model.Student, error) {
	return model.Student{}, nil
}
