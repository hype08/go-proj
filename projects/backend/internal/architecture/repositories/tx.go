package repositories

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TxManager interface {
	Get(txID uuid.UUID) (*sqlx.Tx, error)
}
