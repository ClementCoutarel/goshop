package config

import "database/sql"

type Config struct {
	DB           *sql.DB
	IsProduction bool
}

func NewConfig(db *sql.DB, isProduction bool) Config {
	return Config{
		DB:           db,
		IsProduction: isProduction,
	}
}
