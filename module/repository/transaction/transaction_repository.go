package transaction

import (
	"context"

	MDtransaction "github.com/arudji1211/ecommerce/module/model/transaction"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetAllByUserID(ctx context.Context, db *gorm.DB, UserID uint) (TransactionOut []MDtransaction.Transaction, err error)
	GetByID(ctx context.Context, db *gorm.DB, ID uint) (TransactionOut []MDtransaction.Transaction, err error)
	GetAll(ctx context.Context, db *gorm.DB) (TransactionOut []MDtransaction.Transaction, err error)
	Create(ctx context.Context, db *gorm.DB, TransactionIn MDtransaction.Transaction) (MDtransaction.Transaction, error)
	UpdateByID(ctx context.Context, db *gorm.DB, ID uint, TransactionIn MDtransaction.Transaction) (err error)
	DeleteByID(ctx context.Context, db *gorm.DB, ID uint) (err error)
	IsTransaction(Db *gorm.DB)
}
