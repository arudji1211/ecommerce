package transaction

import (
	"context"

	MDtransaction "github.com/arudji1211/ecommerce/module/model/transaction"
)

type TransactionRepository interface {
	GetAllByUserID(ctx context.Context, UserID uint) (TransactionOut []MDtransaction.Transaction, err error)
	GetByID(ctx context.Context, ID uint) (TransactionOut []MDtransaction.Transaction, err error)
	GetAll(ctx context.Context, UserID uint) (TransactionOut []MDtransaction.Transaction, err error)
	Create(ctx context.Context, ID uint, TransactionIn MDtransaction.Transaction) (MDtransaction.Transaction, error)
	UpdateByID(ctx context.Context, ID uint, TransactionIn MDtransaction.Transaction) (err error)
	DeleteByID(ctx context.Context, ID uint) (err error)
}
