package store

import "github.com/ekediala/expensix/sqlx"

type Store struct {
	db sqlx.DBTX
}

func NewStore(db sqlx.DBTX) *Store {
	return &Store{db}
}
