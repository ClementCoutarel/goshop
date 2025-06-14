package database

import "database/sql"

type ProductSQLRepo struct {
	DB *sql.DB
}

func (r *ProductSQLRepo) GetAll()    {}
func (r *ProductSQLRepo) GetById()   {}
func (r *ProductSQLRepo) Create()    {}
func (r *ProductSQLRepo) UpdateOne() {}
func (r *ProductSQLRepo) DeleteOne() {}
