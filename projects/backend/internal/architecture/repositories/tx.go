package repositories

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type TxManager interface {
	Get(ctx context.Context) (*sqlx.Tx, error)
}
