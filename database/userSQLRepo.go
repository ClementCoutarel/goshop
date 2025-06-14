package database

import (
	"database/sql"
)

type UserSQLRepo struct {
	DB *sql.DB
}

func (r *UserSQLRepo) GetAll()    {}
func (r *UserSQLRepo) GetById()   {}
func (r *UserSQLRepo) Create()    {}
func (r *UserSQLRepo) DeleteOne() {}
func (r *UserSQLRepo) UpdateOne() {}
