package company

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

func (s svc) getAll() ([]model.Company, error) {
	return nil, nil
}

func (s svc) getById(id string) (model.Company, error) {
	return model.Company{}, nil
}

func (s svc) create(c model.Company) (model.Company, error) {
	return model.Company{}, nil
}

func (s svc) update(id string) (model.Company, error) {
	return model.Company{}, nil
}

func (s svc) delete(id string) (model.Company, error) {
	return model.Company{}, nil
}
