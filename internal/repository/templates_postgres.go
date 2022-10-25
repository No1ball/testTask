package repository

import "github.com/jmoiron/sqlx"

type TemplatesPostgres struct {
	db *sqlx.DB
}

func NewTemplatesPostgres(db *sqlx.DB) *TemplatesPostgres {
	return &TemplatesPostgres{db: db}
}
