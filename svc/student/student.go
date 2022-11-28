package student

import (
	"database/sql"
	"github.com/abishekrv-zs/final-assignment/model"
)

type svc struct {
	db *sql.DB
}

func New(db *sql.DB) svc {
	return svc{db: db}
}

func (s svc) getAll() ([]model.Student, error) {
	return nil, nil
}

func (s svc) getById(id string) (model.Student, error) {
	return model.Student{}, nil
}

func (s svc) create(student model.Student) (model.Student, error) {
	return model.Student{}, nil
}

func (s svc) update(id string) (model.Student, error) {
	return model.Student{}, nil
}

func (s svc) delete(id string) (model.Student, error) {
	return model.Student{}, nil
}
