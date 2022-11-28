package student

import (
	"github.com/abishekrv-zs/final-assignment/model"
	"github.com/abishekrv-zs/final-assignment/store"
)

type svc struct {
	store store.StudentStore
}

func New(store store.StudentStore) svc {
	return svc{store: store}
}

func (s svc) GetAll() ([]model.Student, error) {
	return nil, nil
}

func (s svc) GetById(id string) (model.Student, error) {
	return model.Student{}, nil
}

func (s svc) Create(student model.Student) (model.Student, error) {
	return model.Student{}, nil
}

func (s svc) Update(id string) (model.Student, error) {
	return model.Student{}, nil
}

func (s svc) Delete(id string) (model.Student, error) {
	return model.Student{}, nil
}
