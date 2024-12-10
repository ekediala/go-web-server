package store

import "github.com/ekediala/template/sqlx"

type Store struct {
	db sqlx.DBTX
}

func NewStore(db sqlx.DBTX) *Store {
	return &Store{db}
}
