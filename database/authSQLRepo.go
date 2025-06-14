package database

import "database/sql"

type AuthSQLRepo struct {
	DB *sql.DB
}

func (r *AuthSQLRepo) Register() {}
func (r *AuthSQLRepo) Signin()   {}
