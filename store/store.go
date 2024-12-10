package store

import "github.com/ekediala/expensix/sqlx"

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db}
}
