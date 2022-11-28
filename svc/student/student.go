package student

import (
	"github.com/abishekrv-zs/final-assignment/model"
	"github.com/abishekrv-zs/final-assignment/store"
)

type Svc struct {
	store store.StudentStore
}

func New(store store.StudentStore) Svc {
	return Svc{store: store}
}

func (s Svc) GetAll() ([]model.Student, error) {
	return nil, nil
}

func (s Svc) GetById(id string) (model.Student, error) {
	return model.Student{}, nil
}

func (s Svc) Create(student model.Student) (model.Student, error) {
	return model.Student{}, nil
}

func (s Svc) Update(id string) (model.Student, error) {
	return model.Student{}, nil
}

func (s Svc) Delete(id string) (model.Student, error) {
	return model.Student{}, nil
}
