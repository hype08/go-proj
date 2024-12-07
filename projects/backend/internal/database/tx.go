package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type txRecord struct {
	label string
	tx    *sqlx.Tx
}

type TxManager struct {
	pool    *sqlx.DB
	records map[string]*txRecord
}

func NewTxManager(db *sqlx.DB) *TxManager {
	return &TxManager{
		pool:    db,
		records: make(map[string]*txRecord),
	}
}

func (tm *TxManager) Begin(label string) error {
	if _, exists := tm.records[label]; exists {
		return fmt.Errorf("transaction with label '%s' already exists", label)
	}

	tx, err := tm.pool.Beginx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	tm.records[label] = &txRecord{
		label: label,
		tx:    tx,
	}
	return nil
}

// Commit commits the transaction with the given label
func (tm *TxManager) Commit(label string) error {
	record, exists := tm.records[label]
	if !exists {
		return fmt.Errorf("no transaction found with label '%s'", label)
	}

	if err := record.tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	delete(tm.records, label)
	return nil
}

// Rollback rolls back the transaction with the given label
func (tm *TxManager) Rollback(label string) error {
	record, exists := tm.records[label]
	if !exists {
		return fmt.Errorf("no transaction found with label '%s'", label)
	}

	if err := record.tx.Rollback(); err != nil {
		return fmt.Errorf("failed to rollback transaction: %w", err)
	}

	delete(tm.records, label)
	return nil
}

// GetTx returns the transaction associated with the given label
func (tm *TxManager) GetTx(label string) (*sqlx.Tx, error) {
	record, exists := tm.records[label]
	if !exists {
		return nil, fmt.Errorf("no transaction found with label '%s'", label)
	}
	return record.tx, nil
}
