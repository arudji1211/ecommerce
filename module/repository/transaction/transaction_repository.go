package transaction

import (
	"context"

	"github.com/arudji1211/ecommerce/module/model/transaction"
)

type TransactionRepository interface {
	GetAllByUserID(ctx context.Context, UserID uint) (TransactionOut []transaction.Transaction, err error)
	GetByID(ctx context.Context, ID uint) (TransactionOut []transaction.Transaction, err error)
	GetAll(ctx context.Context, UserID uint) (TransactionOut []transaction.Transaction, err error)
	UpdateByID(ctx context.Context, ID uint, TransactionIn transaction.Transaction) (err error)
	DeleteByID(ctx context.Context, ID uint) (err error)
}
