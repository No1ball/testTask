package repository

import "github.com/jmoiron/sqlx"

type Followers interface {
}
type Templates interface {
}
type Api interface {
}
type Repository struct {
	Followers
	Templates
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Followers: NewFollowersPostgres(db),
		Templates: NewTemplatesPostgres(db),
	}
}
