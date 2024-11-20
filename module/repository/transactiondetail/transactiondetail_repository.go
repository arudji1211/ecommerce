package transactiondetail

import (
	"context"

	MDtransactiondetail "github.com/arudji1211/ecommerce/module/model/transactionDetail"
	"gorm.io/gorm"
)

type TransactionDetailRepository interface {
	GetAllByTransactionID(ctx context.Context, db *gorm.DB, TransactionID uint) (TransactionDetailOut []MDtransactiondetail.TransactionDetail, err error)
	GetByTransactionDetailID(ctx context.Context, db *gorm.DB, TransactionID uint) (TransactionDetailOut MDtransactiondetail.TransactionDetail, err error)
	Create(ctx context.Context, db *gorm.DB, TransactionDetailIn MDtransactiondetail.TransactionDetail) (MDtransactiondetail.TransactionDetail, error)
	UpdateByTransactionDetailID(ctx context.Context, db *gorm.DB, TransactionDetailIn MDtransactiondetail.TransactionDetail) (err error)
	DeleteByTransactionDetailID(ctx context.Context, db *gorm.DB, TransactionID uint) (err error)
	IsTransaction(Db *gorm.DB)
}
