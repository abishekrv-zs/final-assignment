package company

import (
	"github.com/abishekrv-zs/final-assignment/model"
	"github.com/abishekrv-zs/final-assignment/store"
)

type Svc struct {
	store store.CompanyStore
}

func New(store store.CompanyStore) Svc {
	return Svc{store: store}
}

func (s Svc) GetAll() ([]model.Company, error) {
	return nil, nil
}

func (s Svc) GetById(id string) (model.Company, error) {
	return model.Company{}, nil
}

func (s Svc) Create(c model.Company) (model.Company, error) {
	return model.Company{}, nil
}

func (s Svc) Update(id string) (model.Company, error) {
	return model.Company{}, nil
}

func (s Svc) Delete(id string) (model.Company, error) {
	return model.Company{}, nil
}
