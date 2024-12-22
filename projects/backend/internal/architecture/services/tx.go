package services

import (
	"context"
)

type TxManager interface {
	Transaction(ctx context.Context, label string, operation func(ctx context.Context) error) error
}
