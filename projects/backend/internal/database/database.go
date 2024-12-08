package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	Pool *sqlx.DB
	Txm  *TxManager
}

func NewDatabase(databaseConnection string) (*DB, error) {
	pool, err := sqlx.Connect("postgres", databaseConnection)

	return &DB{
		Pool: pool,
	}, err
}
