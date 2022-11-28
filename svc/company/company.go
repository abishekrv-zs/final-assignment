package company

import (
	"github.com/abishekrv-zs/final-assignment/model"
	"github.com/abishekrv-zs/final-assignment/store"
)

type svc struct {
	store store.CompanyStore
}

func New(store store.CompanyStore) svc {
	return svc{store: store}
}

func (s svc) GetAll() ([]model.Company, error) {
	return nil, nil
}

func (s svc) GetById(id string) (model.Company, error) {
	return model.Company{}, nil
}

func (s svc) Create(c model.Company) (model.Company, error) {
	return model.Company{}, nil
}

func (s svc) Update(id string) (model.Company, error) {
	return model.Company{}, nil
}

func (s svc) Delete(id string) (model.Company, error) {
	return model.Company{}, nil
}
