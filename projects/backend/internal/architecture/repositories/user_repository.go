package repositories

import (
	"log"
)

type UserRepository interface{}

type userRepository struct {
	txm TxManager
}

func NewUserRepository(txm TxManager) UserRepository {
	if txm == nil {
		log.Fatal("Transaction manager nil pointer")
	}

	return &userRepository{
		txm: txm,
	}
}
