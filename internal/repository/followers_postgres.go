package repository

import "github.com/jmoiron/sqlx"

type FollowersPostgres struct {
	db *sqlx.DB
}

func NewFollowersPostgres(db *sqlx.DB) *FollowersPostgres {
	return &FollowersPostgres{
		db: db,
	}
}
