package pgtxmanager

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// TxManager represents a transaction manager
//
//go:generate minimock -o mock/ -s "_minimock.go"
type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

// Handler represents a handler for a transaction
type Handler func(ctx context.Context) error

// Transactor represents a transactor
//
//go:generate minimock -o mock/ -s "_minimock.go"
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}
